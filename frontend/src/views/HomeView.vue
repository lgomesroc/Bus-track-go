<script>
import BusList from '../components/BusList.vue'

export default {
  components: {
    BusList
  },

  data() {
    return {
      buses: [],
      loading: true,
      error: null
    }
  },

  mounted() {
    fetch('http://localhost:8080/api/buses')
      .then(response => {
        if (!response.ok) {
          throw new Error('Falha ao carregar os ônibus')
        }

        return response.json()
      })
      .then(data => {
        this.buses = data
      })
      .catch(error => {
        this.error = error.message
      })
      .finally(() => {
        this.loading = false
      })
  }
}
</script>

<template>
  <main>
    <h1>BusTrack</h1>

    <h2>Ônibus cadastrados</h2>

    <p v-if="loading">Carregando ônibus...</p>

    <p v-else-if="error">{{ error }}</p>

    <BusList v-else :buses="buses" />
  </main>
</template>