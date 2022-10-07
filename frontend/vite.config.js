import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import Unocss from 'unocss/vite'
import { presetUno, presetAttributify } from 'unocss'
import path from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), Unocss({
    presets: [
      presetUno(),
      presetAttributify()
    ],
    shortcuts: {
      'btn': 'py-2 px-4 font-semibold rounded-lg shadow-md cursor-pointer',
      'btn-green': 'text-white bg-green-500 hover:bg-green-700',
      // single utility alias
      'red': 'text-red-100'
    }
  })],
  resolve:{
    alias:{
      '~': path.resolve(__dirname),
      '@': path.resolve(__dirname, 'src')
    }
  }
})
