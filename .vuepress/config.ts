import { defineUserConfig } from '@vuepress/cli'
import type { ThemeOptions } from './theme/shared'
import { path } from '@vuepress/utils'
import type { ViteBundlerOptions } from '@vuepress/bundler-vite'
import vueJsx from '@vitejs/plugin-vue-jsx'

export default defineUserConfig<ThemeOptions, ViteBundlerOptions>({
    lang: 'zh-CN',
    title: '代码炼金工坊',
    description: 'Go 项目结构实践',

    theme: path.resolve(__dirname, 'theme/index.tsx'),
    themeConfig: {
      title: "Go 项目结构实践",
      name: "科学捜査官",
      avatar: "./yuchanns.jpg",
      socials: {
        github: "yuchanns",
        twitter: "realyuchanns",
      },
      desc: "Enjoy Go/Rust. Loving Anime Girls. Vim User. Fan of LiSA.",
      description: "追寻计算机炼金术的贤者之石",
      copyright: "<a href=\"https://beian.miit.gov.cn/\">闽ICP备2020021086号-1</a>",
      startDate: 2021,
      nav: [
        { name: "About", link: "https://yuchanns.xyz/" }
      ]
    },

    bundler: '@vuepress/bundler-vite',
    bundlerConfig: {
      viteOptions: {
        plugins: [
          vueJsx({})
        ]
      }
    },

    markdown: {
      code: {
        lineNumbers: false
      }
    }
})
