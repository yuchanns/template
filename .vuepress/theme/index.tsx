const { path } = require('@vuepress/utils')

const theme = (opts, app) => {
  return {
      name: 'vuepress-theme-yuchanns',
      layouts: {
          Layout: path.resolve(__dirname, 'layouts/layout.tsx'),
          404: path.resolve(__dirname, 'layouts/notfound.tsx')
      },
      plugins: [
        [
          '@vuepress/plugin-theme-data',
          { themeData: opts }
        ]
      ]
  }
}

module.exports = theme
