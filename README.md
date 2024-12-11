<p align="right">
   <strong>中文</strong> | <a href="./README.en.md">English</a>
</p>

<p align="center">
   <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://github.com/MartialBE/one-hub/assets/42402987/c4125d1a-5577-446d-ba15-2a71c52140c1">
   <img height="90" src="https://raw.githubusercontent.com/MartialBE/one-api/main/web/src/assets/images/logo.svg">
   </picture>
</p>

<div align="center">

# One Hub

_本项目是基于[one-api](https://github.com/songquanpeng/one-api)二次开发的[one-hub](https://github.com/MartialBE/one-hub)再次开发而来的_

<p align="center">
  <a href="https://raw.githubusercontent.com/MartialBE/one-api/main/LICENSE">
    <img src="https://img.shields.io/github/license/MartialBE/one-api?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/MartialBE/one-hub/releases/latest">
    <img src="https://img.shields.io/github/v/release/MartialBE/one-api?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://github.com/users/MartialBE/packages/container/package/one-api">
    <img src="https://img.shields.io/badge/docker-ghcr.io-blue" alt="docker">
  </a>
  <a href="https://hub.docker.com/r/martialbe/one-api">
    <img src="https://img.shields.io/badge/docker-dockerHub-blue" alt="docker">
  </a>
  <a href="https://goreportcard.com/report/github.com/MartialBE/one-api">
    <img src="https://goreportcard.com/badge/github.com/MartialBE/one-api" alt="GoReportCard">
  </a>
</p>

**请不要和原版混用，因为新增功能，数据库与原版不兼容**

**为了更加简洁，本项目之后，除了新增供应商时会更新程序自带的模型列表，平常不再更新程序自带的模型列表。**

**如果发现缺少新模型，请在`后台-模型价格-更新价格`中更新新增的模型**

[演示网站](https://one-api-martialbe.vercel.app/)

</div>

> [!WARNING]
> 本项目为个人学习使用，不保证稳定性，且不提供任何技术支持，使用者必须在遵循 OpenAI 的使用条款以及法律法规的情况下使用，不得用于非法用途。  
> 根据[《生成式人工智能服务管理暂行办法》](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm)的要求，请勿对中国地区公众提供一切未经备案的生成式人工智能服务。

## 功能变化

- 重构了claude的chat.go文件，实现了提示缓存 （当前支持system prompt，大于2000字符时，会自动缓存，或者自行填入中文符号、、符号后的内容会进行缓存）
- docker compose文件修改为本地代码构建，方便开发使用

## 文档

请查看[文档](https://github.com/MartialBE/one-hub/wiki)

## 当前支持的供应商

| 供应商                                                                | Chat                     | Embeddings | Audio  | Images      | 其他                                                             |
| --------------------------------------------------------------------- | ------------------------ | ---------- | ------ | ----------- | ---------------------------------------------------------------- |
| [OpenAI](https://platform.openai.com/docs/api-reference/introduction) | ✅                       | ✅         | ✅     | ✅          | -                                                                |
| [Azure OpenAI](https://oai.azure.com/)                                | ✅                       | ✅         | ✅     | ✅          | -                                                                |
| [Azure Speech](https://portal.azure.com/)                             | -                        | -          | ⚠️ tts | -           | -                                                                |
| [Anthropic](https://www.anthropic.com/)                               | ✅                       | -          | -      | -           | -                                                                |
| [Gemini](https://aistudio.google.com/)                                | ✅                       | -          | -      | -           | -                                                                |
| [百度文心](https://console.bce.baidu.com/qianfan/overview)            | ✅                       | ✅         | -      | -           | -                                                                |
| [通义千问](https://dashscope.console.aliyun.com/overview)             | ✅                       | ✅         | -      | -           | -                                                                |
| [讯飞星火](https://console.xfyun.cn/)                                 | ✅                       | -          | -      | -           | -                                                                |
| [智谱](https://open.bigmodel.cn/overview)                             | ✅                       | ✅         | -      | ⚠️ 图片生成 | -                                                                |
| [腾讯混元](https://cloud.tencent.com/product/hunyuan)                 | ✅                       | -          | -      | -           | -                                                                |
| [百川](https://platform.baichuan-ai.com/console/apikey)               | ✅                       | ✅         | -      | -           | -                                                                |
| [MiniMax](https://www.minimaxi.com/user-center/basic-information)     | ✅                       | ✅         | -      | -           | -                                                                |
| [Deepseek](https://platform.deepseek.com/usage)                       | ✅                       | -          | -      | -           | -                                                                |
| [Moonshot](https://moonshot.ai/)                                      | ✅                       | -          | -      | -           | -                                                                |
| [Mistral](https://mistral.ai/)                                        | ✅                       | ✅         | -      | -           | -                                                                |
| [Groq](https://console.groq.com/keys)                                 | ✅                       | -          | -      | -           | -                                                                |
| [Amazon Bedrock](https://console.aws.amazon.com/bedrock/home)         | ⚠️ 仅支持 Anthropic 模型 | -          | -      | -           | -                                                                |
| [零一万物](https://platform.lingyiwanwu.com/details)                  | ✅                       | -          | -      | -           | -                                                                |
| [Cloudflare AI](https://ai.cloudflare.com/)                           | ✅                       | -          | ⚠️ stt | ⚠️ 图片生成 | -                                                                |
| [Midjourney](https://www.midjourney.com/)                             | -                        | -          | -      | -           | [midjourney-proxy](https://github.com/novicezk/midjourney-proxy) |
| [Cohere](https://cohere.com/)                                         | ✅                       | -          | -      | -           | -                                                                |
| [Stability AI](https://platform.stability.ai/account/credits)         | -                        | -          | -      | ⚠️ 图片生成 | -                                                                |
| [Coze](https://www.coze.com/open/docs/chat?_lang=zh)                  | ✅                       | -          | -      | -           | -                                                                |
| [Ollama](https://github.com/ollama/ollama)                            | ✅                       | ✅         | -      | -           | -                                                                |
| [Suno](https://suno.com/)                                             | -                        | -          | -      | -           | [Suno-API](https://github.com/Suno-API/Suno-API)                 |


## 交流群

<img src="https://github.com/MartialBE/one-hub/assets/42402987/9b608d39-70ae-4b2e-be49-09afab6bd536" width="300">

