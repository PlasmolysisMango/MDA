<!-- markdownlint-disable MD033 MD041 -->
<p align="center">
  <img alt="LOGO" src="https://s1.imagehub.cc/images/2026/05/12/d1d0730a19f251d8ea800897754f0ab2.png" width="256" height="256" />
</p>

<div align="center">

# MDA

Maa Doro Assistant

**[简体中文](README.md)** | **English**

</div>

MDA is a game automation assistant built on [MaaFramework](https://github.com/MaaXYZ/MaaFramework), rewritten from [DoroHelper](https://github.com/1204244136/DoroHelper).

---

## ⚠️ Account Ban Risk

Some users have recently reported that their game accounts were banned after using this software; some have also reported that their accounts have since been unbanned, so those bans may have included official mistakes or false positives from risk controls. MDA is a third-party automation assistant, and using it may still violate the game's terms of service or risk-control rules.

**Please use this software only after fully understanding and accepting the risks, including account bans and data loss.**

---

## Language Compatibility

MDA's interface supports multiple languages including Chinese and English, but **the script's functionality is currently only adapted for the Chinese game interface**.

If you are using an English or other language game interface, you may encounter recognition errors or functional issues. If you experience errors, please switch your game to **Simplified Chinese** first and try again. If the problem persists after switching, feel free to submit feedback and I'll help investigate.

---

## Getting Started

### 1. First Launch

Take a moment to explore the interface before running any tasks to understand the available features and settings.

### 2. Set Up Hotkeys (Recommended)

Go to **Settings (top-right corner) → Hotkeys** and enable global hotkeys, in case the program becomes unresponsive and you need to exit.

---

## Reporting Issues

If the script encounters an error, follow these steps to collect information for troubleshooting.

### Step 1: Enable Debug Images

1. Go to **Settings → Debug**
2. Enable **Save Debug Images**

> ⚠️ This option must be re-enabled each time you start the program.

### Step 2: Reproduce the Issue

Debug mode saves a screenshot for every action, so **avoid running tasks for extended periods** — it will generate a large number of images and consume disk space.

Recommended approach:

- After enabling debug mode, **only run the task that has the issue**
- Stop immediately after reproducing the problem and prepare to package the logs

### Step 3: Export the Logs

1. Click the **Export Logs** icon next to **Run Log** in the bottom-right corner
2. In the `debug` folder that opens, find the generated archive
3. Send the archive to the developer

> 💡 After submitting feedback, it's recommended to **delete the old `vision` folder** and restart the program. This keeps debug images from different issues separate and makes troubleshooting easier.
