[![github-workflow-shield]][github-workflow-url]
[![Contributors][contributors-shield]][contributors-url]
[![MIT License][license-shield]][license-url]
[![Go Version][gomod]][gomod-url]

<br />

<div>
<h3 align="center">go cli template</h3>

  <p align="center">
    Project template for a Golang CLI project.
  </p>
</div>

## ℹ️ About

This project contains a template for a golang based cli application.

### What is included

- [Cobra](https://cobra.dev/) - A Framework for Modern CLI Apps in Go
- [GoReleaser](https://goreleaser.com/) - Used to create a release of the application
- Linting with [golangci-lint](https://golangci-lint.run/)
- [Task](https://taskfile.dev/) - Task runner.
- [lefthook](https://github.com/evilmartians/lefthook) - Fast and powerful Git hooks manager for any type of projects.
- [Devbox](https://www.jetpack.io/devbox) - Isolated development envrionment, powered by Nix.
- [GitHub Actions](https://github.com/actions) - CI/CD Pipeline

### Built With

* Golang

## 🚀 Getting Started

Start by installing gonew using go install:

```shell
go install golang.org/x/tools/cmd/gonew@latest
```

To copy an existing template, run gonew in your new project’s parent directory with two arguments: first, the path to the template you wish to copy, and second, the module name of the project you are creating. For example:

```shell
gonew github.com/brpaz/go-cli-template example.com/my-cli
```

## 🤝 Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 🫶 Support

If you find this project helpful and would like to support its development, there are a few ways you can contribute:

[![Sponsor me on GitHub](https://img.shields.io/badge/Sponsor-%E2%9D%A4-%23db61a2.svg?&logo=github&logoColor=red&&style=for-the-badge&labelColor=white)](https://github.com/sponsors/brpaz)

<a href="https://www.buymeacoffee.com/Z1Bu6asGV" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## 📃 License

Distributed under the MIT License. See [LICENSE](LICENSE.md) file for details.

## 📩 Contact

- Bruno Paz - [https://brunopaz.dev](https://brunopaz.dev) - oss@brunopaz.dev

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/brpaz/go-cli-template.svg?style=for-the-badge
[contributors-url]: https://github.com/brpaz/go-cli-template/graphs/contributors
[license-shield]: https://img.shields.io/github/license/brpaz/go-cli-template?style=for-the-badge
[license-url]: https://github.com/brpaz/go-cli-template/LICENSE
[github-workflow-shield]: https://img.shields.io/github/actions/workflow/status/brpaz/go-cli-template/ci.yml?style=for-the-badge
[github-workflow-url]: https://github.com/brpaz/go-cli-template/actions
[gomod]: https://github.com/brpaz/go-cli-template
[gomod-url]: https://img.shields.io/github/go-mod/go-version/brpaz/go-cli-template
