import { defineComponent } from 'vue'
import ToggleTheme from './toggleTheme'
import { useThemeLocaleData } from '../shared'

export default defineComponent({
  setup() {
    const data = useThemeLocaleData().value
    return () => (
      <header class="header">
        <h4 class="site-title typography-title">
          <a href="/" class="site-title-link">{ data.title }</a>
        </h4>
        <div class="spacer"></div>
        <div class="nav">
          { data.nav.map(item => <a href={ item.link }>{ item.name }</a>) }
          <ToggleTheme />
        </div>
      </header>
    )
  }
})

