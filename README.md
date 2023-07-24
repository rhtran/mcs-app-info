<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links

[![NPM Version][npm-image]][npm-url]
[![Build Status][travis-image]][travis-url]
[![Downloads Stats][npm-downloads]][npm-url]


[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
-->

<!-- PROJECT LOGO -->


<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#api-endpoints">API Endpoints</a></li>
    <li><a href="./docs/project-layout.md">Project Layout</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>


# mcs-app-info

## Publish Library
1. git tag "vx.x.x"
2. git push --tag


## [go get private repos using SSH key auth instead of password auth](https://gist.github.com/StevenACoffman/866b06ed943394fbacb60a45db5982f2)
1. add export GOPRIVATE=github.com/{{your username}} 
2. add export GITHUB_TOKEN=your-access-token-generated-from-github
3. git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
4. cat .gitconfig

Look something like this
```text
[url "https://your-access-token-generated-from-github:x-oauth-basic@github.com/"]
        insteadOf = https://github.com/

```