import Aura from '@primeuix/themes/aura';
import { definePreset } from '@primeuix/themes';

const a5Preset = definePreset(Aura, {
  semantic: {
    primary: {
      50: '{blue.50}',
      100: '{blue.100}',
      200: '{blue.200}',
      300: '{blue.300}',
      400: '{blue.400}',
      500: '{blue.500}',
      600: '{blue.600}',
      700: '{blue.700}',
      800: '{blue.800}',
      900: '{blue.900}',
      950: '{blue.950}',
    }
  }
});

export default defineNuxtConfig({
  compatibilityDate: '2025-05-15',
  devtools: { enabled: true },
  modules: [
    '@nuxt/eslint',
    '@nuxt/fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxt/scripts',
    '@nuxt/test-utils',
    '@primevue/nuxt-module',
  ],
  primevue: {
    options: {
      theme: {
        preset: a5Preset,
        options: {
          darkModeSelector: 'system',
          cssLayer: false,
        },
      },
    },
  },
//  css: ['~/assets/css/main.css'],
})
