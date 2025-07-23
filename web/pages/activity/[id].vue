<template>
  <div>
    <h1>{{ route.params.id }}: Start Date&ndash;End Date</h1>
    <Tabs value="0">
      <TabList>
        <Tab value="0">Plan</Tab>
        <Tab value="1">Org Chart</Tab>
        <Tab value="2">Billeting</Tab>
        <Tab value="3">Budget</Tab>
      </TabList>
    </Tabs>
  </div>
</template>

<script setup lang="ts">
const route = useRoute()
</script>

<style scoped>

</style>