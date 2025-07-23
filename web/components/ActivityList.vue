<template>
<DataTable :value="activities"> <!--table-style="min-width: 50rem">-->
  <Column header="">
    <template #body="slotProps">
      <Button>
        <NuxtLink :to="`/activity/${slotProps.data.id}`">
          View
        </NuxtLink>
      </Button>
    </template>
  </Column>
  <Column field="name" header="Activity Name"/>
  <Column field="unit" header="Unit"/>
  <Column field="type" header="Activity Type"/>
  <Column field="start" header="Start Date"/>
  <Column field="end" header="End Date"/>
</DataTable>
</template>

<script setup lang="ts">
const activities = [
  {
    id: "dhxxii",
    unit: "RMR-UT-001",
    name: "Desert Hawk XXII",
    type: "Encampment",
    start: "25 June 2025",
    end: "3 July 2025",
  }
]

</script>

<style scoped>

</style>