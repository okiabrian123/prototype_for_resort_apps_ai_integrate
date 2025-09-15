import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
import dotenv from 'dotenv'
import { resolve } from 'path'

// Load environment variables from root .env file
const env = dotenv.config({ path: resolve(__dirname, '../.env') }).parsed || {}

// Get the domain name from environment variables or use default
const allowedDomain = env.DOMAIN_NAME || 'prototype-resort-apps.okiabrian.my.id'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    host: true,
    allowedHosts: [
      'localhost',
      '127.0.0.1',
      allowedDomain
    ]
  },
  preview: {
    host: true,
    allowedHosts: [
      'localhost',
      '127.0.0.1',
      allowedDomain
    ]
  }
})