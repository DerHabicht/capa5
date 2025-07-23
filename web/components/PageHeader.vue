<template>
    <Menubar :model="items">
      <template #start>
        <Button><NuxtLink to="/">A5</NuxtLink></Button>
      </template>
      <template #item="{ item }">
        <NuxtLink active-class="bg-primary" :to="item.route">{{ item.label }}</NuxtLink>
      </template>
      <template #end>
        <Button><NuxtLink to="/about">About</NuxtLink></Button>
        <ToggleSwitch name="dark"/>
        <Avatar/>
      </template>
    </Menubar>
</template>

<script setup lang="ts">
import type { MenuItem } from "primevue/menuitem";

const items : MenuItem[] = [
  {
    label: "Plan Library",
    route: "/plan",
  },
  {
    label: "Activities",
    route: "/activity",
  },
]
</script>

<style scoped>
</style>