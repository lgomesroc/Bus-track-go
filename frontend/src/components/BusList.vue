<script>
export default {
  props: {
    buses: {
      type: Array,
      required: true
    }
  }
}
</script>

<template>
  <ul class="bus-list">
    <li v-for="bus in buses" :key="bus.id" class="bus-item">
      <strong>{{ bus.prefix }} - {{ bus.model }}</strong>

      <span>Placa: {{ bus.licensePlate }}</span>

      <span>Capacidade: {{ bus.capacity }} passageiros</span>

      <span>Status: {{ bus.status }}</span>
    </li>
  </ul>
</template>

<style scoped>
.bus-list {
  list-style: none;
  padding: 0;
}

.bus-item {
  margin-bottom: 16px;
  padding: 16px;
  border: 1px solid #ddd;
  border-radius: 8px;
}

.bus-item strong {
  display: block;
  margin-bottom: 8px;
}

.bus-item span {
  display: block;
  margin-bottom: 4px;
}
</style>