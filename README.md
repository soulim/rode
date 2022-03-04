# Rode

> Rode – the line stretched between the anchor and the boat.

## Usage

### List podcast episodes

```
rode list <FEED_FILE_PATH>
```

Rode lists all episodes from given RSS file.

```ShellSession
> rode list feed.xml
44dc75b3-534b-4a42-b9cc-e4c16a38108e Building an effective development environment
87024762-154b-4442-8ea2-2a9a39b5cd06 From Rocks to Code: How a Geologist Became a Software Developer
a4a7f989-acb4-4edc-b7db-319a65cfe16b SEO (Search Engine Optimization) and software development
b09359ab-dc41-4d5b-9029-696261fd02ba Remote work
35ea3481-14ed-4e1a-9a9f-6c5e73ba1e7c Experiments
add31084-80e9-48cc-9060-017722fe20a0 Productivity methods and hacks
23db49d5-c078-4ab1-b2ee-462e54519747 Vim for life?
320bf466-e8b7-4011-a849-ae932bb0c96d Heroes and mentors
```

### Export an episode

```
rode export <FEED_FILE_PATH> <GUID>
```

Rode takes a GUID of the episode and outputs a Markdown file suitable for [Hugo](https://gohugo.io/).

```ShellSession
> rode export internal/rode/testdata/feed.xml 44dc75b3-534b-4a42-b9cc-e4c16a38108e
+++
title = "Building an effective development environment"
date = "2022-01-18T18:00:38Z"
draft = false
description = "TODO: Add description"
cover = "TODO: Add cover"
+++

As software engineers, we rely on a lot of tools: editors,  IDEs, linters, version control, terminal emulators, virtual machines...  Tools we buy, download for free, or build ourselves.

They define – or at least affect – the quality of what we produce. They boost our delivery speed and automate boring tasks but sometimes, we must admit, they also drive us crazy. What is clear is that, by using tools every single day, we have come to depend on them.

In this episode of Code && Beyond, we talk about our favourite tools and try to define how we would build the most effective development environment with them.

Starring a special guest: [Juan Ibiapina](https://github.com/juanibiapina), who has built one of the most advanced development environments.

Notes:

- [Alacritty](https://github.com/alacritty/alacritty) - a fast, cross-platform, OpenGL terminal emulator. 

- [NixOS](https://nixos.org/) is based on Nix, a purely functional package management system. 

- [fzf](https://github.com/junegunn/fzf) is a general-purpose command-line fuzzy finder. 

- [Juan's dotfiles.](https://github.com/juanibiapina/dotfiles/) - [null-ls.nvim](https://github.com/jose-elias-alvarez/null-ls.nvim)

- [Starship](https://starship.rs/) - the minimal, blazing-fast, and infinitely customizable prompt for any shell! 

Have any feedback? Send us an email at [codeandbeyond@protonmail.com](mailto:codeandbeyond@protonmail.com) or leave a voice message on [Anchor](https://anchor.fm/codeandbeyond).

---


_Music by Twisterium from _[_Pixabay_](https://pixabay.com/users/twisterium-20030970/)_._
```

### How to use `rode` together with `fzf`

Together with `fzf` `rode` could be used to export one particular episode. The episode is selected using fuzzy-search interface `fzf` provides.

```ShellSession
> rode list internal/rode/testdata/feed.xml \
  | fzf \
  | cut -d " " -f 1 \
  | xargs bin/rode export internal/rode/testdata/feed.xml
```

## Contributing

PRs accepted.

## License

See [LICENSE](LICENSE).
