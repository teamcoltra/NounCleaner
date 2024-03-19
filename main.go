package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var baseDirFlag, iconsDirFlag, outDir string
	var attBool, uidBool bool

	flag.StringVar(&baseDirFlag, "b", "", "Base directory to scan for icons")
	flag.StringVar(&iconsDirFlag, "i", "", "Specific directory within the base directory containing SVG icons")
	flag.StringVar(&outDir, "o", "dist", "Output subdirectory to save cleaned icons")
	flag.BoolVar(&attBool, "a", true, "Enable or disable writing attribution details to a file")
	flag.BoolVar(&uidBool, "u", true, "Enable or disabled keeping the UID of the filename")

	flag.Parse()
	args := flag.Args()

	// Determine base directory
	baseDir := baseDirFlag
	if baseDir == "" {
		if len(args) > 0 {
			baseDir = args[0]
		} else {
			baseDir = "." // Use current directory as default
		}
	}

	// Resolve the absolute path of the base directory
	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		fmt.Printf("Error resolving absolute path of base directory: %v\n", err)
		return
	}

	// Determine icons directory
	var iconsFullPath string
	if iconsDirFlag != "" {
		// If -i is specified, use it as the path for icons
		iconsFullPath = filepath.Join(absBaseDir, iconsDirFlag)
	} else {
		// Use the base directory itself for icons if -i is not provided
		iconsFullPath = absBaseDir
	}

	// Resolve the output directory
	outFullPath := filepath.Join(absBaseDir, outDir)
	if err := os.MkdirAll(outFullPath, os.ModePerm); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	files, err := os.ReadDir(iconsFullPath)
	if err != nil {
		fmt.Println("Error reading icons directory:", err)
		return
	}

	var attributionDetails string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		if filepath.Ext(fileName) != ".svg" {
			continue
		}

		content, err := os.ReadFile(filepath.Join(iconsFullPath, fileName))
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", fileName, err)
			continue
		}

		attributionNote, cleanedContent := cleanSVG(string(content), fileName)

		if attBool {
			attributionDetails += attributionNote + "\n"
		}

		// Remove prefix and write the cleaned file
		newFileName := removePrefix(fileName, uidBool)
		newFilePath := filepath.Join(outFullPath, newFileName)
		if err := os.WriteFile(newFilePath, []byte(cleanedContent), fs.ModePerm); err != nil {
			fmt.Printf("Error writing cleaned file %s: %v\n", newFileName, err)
			continue
		}
		fmt.Printf("Cleaned file written: %s\n", newFileName)
	}

	if attBool && attributionDetails != "" {
		attributionPath := filepath.Join(absBaseDir, "attribution.txt")
		if err := os.WriteFile(attributionPath, []byte(attributionDetails), fs.ModePerm); err != nil {
			fmt.Printf("Error writing attribution details: %v\n", err)
			return
		}
		fmt.Println("Attribution details written to", attributionPath)
	}
}

func cleanSVG(content, fileName string) (string, string) {
	attributionNote := fileName + ": "
	textRegex := regexp.MustCompile(`<text.*?>(.*?)<\/text>`)

	matches := textRegex.FindAllStringSubmatch(content, -1)
	if len(matches) == 0 {
		attributionNote += "No attribution detail"
	} else {
		for _, match := range matches {
			attributionNote += match[1] + " "
		}
	}

	newContent := textRegex.ReplaceAllString(content, "")
	newContent = strings.Replace(newContent, "<svg", `<svg preserveAspectRatio="xMidYMid meet" `, 1)
	newContent = regexp.MustCompile(`viewBox=".*?"`).ReplaceAllString(newContent, `viewBox="0 0 100 100"`)

	return strings.TrimSpace(attributionNote), newContent
}

// removePrefix removes the prefix from the filename, leaving only the main part.
func removePrefix(fileName string, uidBool bool) string {
	parts := strings.SplitN(fileName, "-", 3) // Assuming the format "noun-<name>-<id>.svg"
	if len(parts) == 3 {
		if uidBool {
			return parts[1] + "-" + parts[2] + ".svg" // Returns "<name>-<uid>.svg"
		} else {
			return parts[1] + ".svg" // Return "<name>.svg"
		}
	}
	return fileName // Return original if not matching expected format
}
