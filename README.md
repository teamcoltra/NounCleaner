# NounCleaner Project

## Overview
NounCleaner is a tool inspired by the icon_cleaner project by Rashan Jibowu ([icon_cleaner GitHub](https://github.com/rashanjibowu/icon_cleaner)). It automates the process of cleaning SVG files obtained from TheNounProject.com by removing attribution tags directly from the SVGs and compiling that attribution information into a separate text file. This simplifies the use of SVG icons while ensuring that the original creators are appropriately credited, making attribution effortless and organized.

## Features
- Automatically scans a specified directory for SVG files.
- Removes attribution text tags from SVG files.
- Compiles all attribution information into a single, easy-to-reference text file.
- Supports custom input and output directories through command-line options.
- Cross-platform compatibility (Windows, Linux, MacOS).

## Installation
NounCleaner binaries can be downloaded from the project's releases page. Alternatively, users with Go installed can build the project directly from source.

### Windows & Linux & Mac
1. Download the appropriate binary for your operating system from the releases page.
2. (Optional) To build from source, ensure you have Go installed and run `go build` in the project directory. This will generate the binary executable.

## Usage
NounCleaner offers several command-line options for flexibility:

- `-b`: Specify the base directory to scan for SVG files. Defaults to the current directory if not specified.
- `-i`: Specify the subdirectory within the base directory containing SVG icons. If not provided, the base directory is used.
- `-o`: Specify the output subdirectory where cleaned icons will be saved. Defaults to "dist" within the base directory.
- `-a`: Enable writing of attribution details to a text file. Enabled by default.

I added the option to disable creating the attribution file because you might have already generated it or are already giving credit in another way (or have become a premium member of The Noun Project) but please do not use this tool to circumvent the Creative Commons license. This isn't me giving you a wink and a nudge, genuinely, people made the icons for free they just ask that you give them credit. Thanks! 

### Examples
1. Cleaning SVGs in the current directory: `nouncleaner`
2. Specifying a custom icons directory: `nouncleaner -i path/to/icons`
3. Specifying both base and icons directories: `nouncleaner -b /path/to/project -i icons`
4. Enabling attribution text file generation: `nouncleaner -a`

### Dropping Folder onto Binary
You can also drag and drop a folder containing SVGs onto the NounCleaner binary. This automatically cleans the SVGs in the dropped folder, using the folder as the base directory.

## Credits
This project is inspired by and gives credit to Rashan Jibowu's icon_cleaner project ([icon_cleaner GitHub](https://github.com/rashanjibowu/icon_cleaner)). It builds upon the original idea by providing additional flexibility and a command-line interface for ease of use across different operating systems.

## Contributing
Contributions are welcome! Whether it's suggesting new features, reporting bugs, or improving documentation, your input helps make NounCleaner better for everyone.

## Thank You
Thank you for using NounCleaner. This tool aims to make working with SVGs from TheNounProject.com easier while respecting and acknowledging the creators' contributions. Have a lovely day and enjoy streamlined SVG management with proper attribution!
