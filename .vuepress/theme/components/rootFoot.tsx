import { defineComponent } from 'vue'
import { useThemeLocaleData } from '../shared'

export default defineComponent({
  setup() {
    const data = useThemeLocaleData().value
    const year = new Date().getFullYear()
    return () => (
      <footer class="footer">
        <span class="typography-thin">
          { data.title } © { data.startDate }-{ year } | <span v-html={ data.copyright }></span> | Powerd by <a href="https://v2.vuepress.vuejs.org">Vuepress</a> <a href="https://github.com/yuchanns/yuchanns">Theme Yuchanns</a>
        </span>
      </footer>
    )
  }
})

