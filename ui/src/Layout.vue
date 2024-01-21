<script setup lang="ts">
import { RouterView } from 'vue-router'
import Top from '@/components/Top.vue'
import Menu from '@/components/Menu.vue'
</script>

<template>
  <Top />
  <Menu />

  <RouterView />
</template>