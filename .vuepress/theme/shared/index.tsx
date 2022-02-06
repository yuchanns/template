import type { ThemeData as _ThemeData, ThemeLocaleDataRef } from '@vuepress/plugin-theme-data/lib/client'
import { useThemeLocaleData as _useThemeLocaleData } from '@vuepress/plugin-theme-data/lib/client'

export interface Nav {
  name: string
  link: string
}

export interface Socials {
  github: string,
  twitter: string
}

export type ThemeOptions = _ThemeData<{
  title: string
  avatar: string
  name: string
  description: string
  desc: string
  copyright: string
  startDate: number
  nav: Nav[],
  socials: Socials
}>

export const useThemeLocaleData = (): ThemeLocaleDataRef<ThemeOptions> => _useThemeLocaleData<ThemeOptions>()
