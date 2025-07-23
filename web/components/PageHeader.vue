<template>
    <Menubar>
      <template #start>
        <Button><NuxtLink to="/">A5</NuxtLink></Button>
      </template>
      <template #end>
        <Button><NuxtLink to="/about">About</NuxtLink></Button>
        <ToggleSwitch name="dark"/>
        <Avatar/>
      </template>
    </Menubar>
</template>

<script setup lang="ts">
</script>
