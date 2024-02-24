# Markdown Quality Checker

The Markdown Quality Checker is a lightweight tool designed to check the quality of your Markdown files. It offers several useful features to help you maintain the quality of your Markdown documents.

## Features

1. **HTTP Links Checking:**
   - This feature checks that all HTTP links in your Markdown files work correctly. It detects broken links and allows you to fix them quickly.

2. **Image Links Checking:**
   - The Markdown Quality Checker also verifies all links to images in your Markdown files. It ensures that these links point to valid and accessible images.

3. **Image Compression Checking:**
   - This feature checks if all images included in your Markdown files are compressed. It helps you reduce the size of your Markdown documents by compressing uncompressed images.

## Installation

To install the Markdown Quality Checker, you need to have Go installed on your system. Once Go is installed, you can use the following command:

```bash
go install github.com/dim-ops/mdqc
```

## Usage

After installing the tool, you can run it using the following command in the directory containing your Markdown files:

```bash
mdqc
```

For advanced usage, you can specify the path to a specific Markdown file or a directory containing multiple Markdown files to check:

```bash
mdqc link -p path
mdqc img -p path -i imgPath
mdqc compress -p imgPath
```

## Contributing

Contributions are welcome! If you have any ideas to improve the Markdown Quality Checker or if you want to add new features, feel free to open a Pull Request.

## License

This project is licensed under the [Apache License](LICENSE).
