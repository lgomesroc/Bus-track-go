<script>
import BusList from '../components/BusList.vue'

export default {
  components: {
    BusList
  },

  data() {
    return {
      buses: [],
      lines: [],
      trips: [],

      loading: true,
      error: null,

      newLine: {
        code: '',
        origin: '',
        destination: ''
      },

      newBus: {
        prefix: '',
        licensePlate: '',
        model: '',
        capacity: 0,
        status: 'active'
      },

      editingBus: null,

      newTrip: {
        lineId: '',
        busId: '',
        tripDate: '',
        tripTime: '',
        passengers: 0
      },

      savingLine: false,
      savingBus: false,
      savingTrip: false,
      deletingBus: false
    }
  },

  mounted() {
    this.loadData()
  },

  methods: {
    async loadData() {
      this.loading = true
      this.error = null

      try {
        const [busesResponse, linesResponse, tripsResponse] =
          await Promise.all([
            fetch('http://localhost:8080/api/buses'),
            fetch('http://localhost:8080/api/lines'),
            fetch('http://localhost:8080/api/trips')
          ])

        if (!busesResponse.ok) {
          throw new Error('Falha ao carregar os ônibus')
        }

        if (!linesResponse.ok) {
          throw new Error('Falha ao carregar as linhas')
        }

        if (!tripsResponse.ok) {
          throw new Error('Falha ao carregar as viagens')
        }

        this.buses = (await busesResponse.json()) || []
        this.lines = (await linesResponse.json()) || []
        this.trips = (await tripsResponse.json()) || []
      } catch (error) {
        this.error = error.message
      } finally {
        this.loading = false
      }
    },

    async createLine() {
      this.savingLine = true
      this.error = null

      try {
        const response = await fetch('http://localhost:8080/api/lines', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(this.newLine)
        })

        if (!response.ok) {
          throw new Error('Falha ao cadastrar a linha')
        }

        const line = await response.json()

        this.lines.push(line)

        this.newLine = {
          code: '',
          origin: '',
          destination: ''
        }
      } catch (error) {
        this.error = error.message
      } finally {
        this.savingLine = false
      }
    },

    async createBus() {
      this.savingBus = true
      this.error = null

      try {
        const response = await fetch('http://localhost:8080/api/buses', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(this.newBus)
        })

        if (!response.ok) {
          throw new Error('Falha ao cadastrar o ônibus')
        }

        const bus = await response.json()

        this.buses.push(bus)

        this.newBus = {
          prefix: '',
          licensePlate: '',
          model: '',
          capacity: 0,
          status: 'active'
        }
      } catch (error) {
        this.error = error.message
      } finally {
        this.savingBus = false
      }
    },

    startEditBus(bus) {
      this.editingBus = {
        ...bus
      }
    },

    cancelEditBus() {
      this.editingBus = null
    },

    async updateBus() {
      if (!this.editingBus) {
        return
      }

      this.savingBus = true
      this.error = null

      try {
        const response = await fetch(
          `http://localhost:8080/api/buses/${this.editingBus.id}`,
          {
            method: 'PUT',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({
              prefix: this.editingBus.prefix,
              licensePlate: this.editingBus.licensePlate,
              model: this.editingBus.model,
              capacity: Number(this.editingBus.capacity),
              status: this.editingBus.status
            })
          }
        )

        if (!response.ok) {
          throw new Error('Falha ao atualizar o ônibus')
        }

        const updatedBus = await response.json()

        const index = this.buses.findIndex(
          bus => bus.id === updatedBus.id
        )

        if (index !== -1) {
          this.buses[index] = updatedBus
        }

        this.editingBus = null
      } catch (error) {
        this.error = error.message
      } finally {
        this.savingBus = false
      }
    },

    async deleteBus(bus) {
      if (!confirm(`Deseja excluir o ônibus ${bus.prefix}?`)) {
        return
      }

      this.deletingBus = true
      this.error = null

      try {
        const response = await fetch(
          `http://localhost:8080/api/buses/${bus.id}`,
          {
            method: 'DELETE'
          }
        )

        if (!response.ok) {
          throw new Error('Falha ao excluir o ônibus')
        }

        this.buses = this.buses.filter(
          item => item.id !== bus.id
        )
      } catch (error) {
        this.error = error.message
      } finally {
        this.deletingBus = false
      }
    },

    async createTrip() {
      this.savingTrip = true
      this.error = null

      try {
        const response = await fetch('http://localhost:8080/api/trips', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            lineId: Number(this.newTrip.lineId),
            busId: Number(this.newTrip.busId),
            tripDate: this.newTrip.tripDate,
            tripTime: this.newTrip.tripTime,
            passengers: Number(this.newTrip.passengers)
          })
        })

        if (!response.ok) {
          throw new Error('Falha ao registrar a viagem')
        }

        const trip = await response.json()

        this.trips.push(trip)

        this.newTrip = {
          lineId: '',
          busId: '',
          tripDate: '',
          tripTime: '',
          passengers: 0
        }
      } catch (error) {
        this.error = error.message
      } finally {
        this.savingTrip = false
      }
    },

    getLineDescription(lineId) {
      const line = this.lines.find(item => item.id === lineId)

      if (!line) {
        return 'Linha não encontrada'
      }

      return `${line.code} - ${line.origin} → ${line.destination}`
    },

    getBusDescription(busId) {
      const bus = this.buses.find(item => item.id === busId)

      if (!bus) {
        return 'Ônibus não encontrado'
      }

      return `${bus.prefix} - ${bus.model}`
    },

    getAveragePassengers(lineId) {
      const lineTrips = this.trips.filter(
        trip => trip.lineId === lineId
      )

      if (lineTrips.length === 0) {
        return 0
      }

      const totalPassengers = lineTrips.reduce(
        (total, trip) => total + trip.passengers,
        0
      )

      return (totalPassengers / lineTrips.length).toFixed(1)
    }
  }
}
</script>

<template>
  <main class="container">
    <h1>BusTrack</h1>

    <p v-if="loading">
      Carregando dados...
    </p>

    <template v-else>
      <p v-if="error" class="error">
        {{ error }}
      </p>

      <section>
        <h2>Cadastrar linha</h2>

        <form @submit.prevent="createLine">
          <input
            v-model="newLine.code"
            type="text"
            placeholder="Código da linha"
            required
          />

          <input
            v-model="newLine.origin"
            type="text"
            placeholder="Origem"
            required
          />

          <input
            v-model="newLine.destination"
            type="text"
            placeholder="Destino"
            required
          />

          <button
            type="submit"
            :disabled="savingLine"
          >
            {{ savingLine ? 'Salvando...' : 'Cadastrar linha' }}
          </button>
        </form>
      </section>

      <section>
        <h2>Linhas cadastradas</h2>

        <p v-if="lines.length === 0">
          Nenhuma linha cadastrada.
        </p>

        <ul v-else>
          <li
            v-for="line in lines"
            :key="line.id"
          >
            <strong>{{ line.code }}</strong>
            - {{ line.origin }} → {{ line.destination }}

            <span>
              | Média de passageiros:
              <strong>
                {{ getAveragePassengers(line.id) }}
              </strong>
            </span>
          </li>
        </ul>
      </section>

      <section>
        <h2>Cadastrar ônibus</h2>

        <form
          v-if="!editingBus"
          @submit.prevent="createBus"
        >
          <input
            v-model="newBus.prefix"
            type="text"
            placeholder="Prefixo"
            required
          />

          <input
            v-model="newBus.licensePlate"
            type="text"
            placeholder="Placa"
            required
          />

          <input
            v-model="newBus.model"
            type="text"
            placeholder="Modelo"
            required
          />

          <input
            v-model.number="newBus.capacity"
            type="number"
            min="1"
            placeholder="Capacidade"
            required
          />

          <select v-model="newBus.status">
            <option value="active">
              Ativo
            </option>

            <option value="maintenance">
              Manutenção
            </option>

            <option value="inactive">
              Inativo
            </option>
          </select>

          <button
            type="submit"
            :disabled="savingBus"
          >
            {{ savingBus ? 'Salvando...' : 'Cadastrar ônibus' }}
          </button>
        </form>

        <form
          v-else
          @submit.prevent="updateBus"
        >
          <input
            v-model="editingBus.prefix"
            type="text"
            placeholder="Prefixo"
            required
          />

          <input
            v-model="editingBus.licensePlate"
            type="text"
            placeholder="Placa"
            required
          />

          <input
            v-model="editingBus.model"
            type="text"
            placeholder="Modelo"
            required
          />

          <input
            v-model.number="editingBus.capacity"
            type="number"
            min="1"
            placeholder="Capacidade"
            required
          />

          <select v-model="editingBus.status">
            <option value="active">
              Ativo
            </option>

            <option value="maintenance">
              Manutenção
            </option>

            <option value="inactive">
              Inativo
            </option>
          </select>

          <button
            type="submit"
            :disabled="savingBus"
          >
            {{ savingBus ? 'Salvando...' : 'Salvar alterações' }}
          </button>

          <button
            type="button"
            @click="cancelEditBus"
          >
            Cancelar
          </button>
        </form>
      </section>

      <section>
        <h2>Ônibus cadastrados</h2>

        <p v-if="buses.length === 0">
          Nenhum ônibus cadastrado.
        </p>

        <ul v-else>
          <li
            v-for="bus in buses"
            :key="bus.id"
          >
            <strong>{{ bus.prefix }}</strong>
            - {{ bus.model }}
            - {{ bus.licensePlate }}
            - Capacidade: {{ bus.capacity }}
            - Status: {{ bus.status }}

            <button
              type="button"
              @click="startEditBus(bus)"
            >
              Editar
            </button>

            <button
              type="button"
              @click="deleteBus(bus)"
              :disabled="deletingBus"
            >
              Excluir
            </button>
          </li>
        </ul>
      </section>

      <section>
        <h2>Registrar viagem</h2>

        <form @submit.prevent="createTrip">
          <select
            v-model="newTrip.lineId"
            required
          >
            <option
              value=""
              disabled
            >
              Selecione a linha
            </option>

            <option
              v-for="line in lines"
              :key="line.id"
              :value="line.id"
            >
              {{ line.code }} -
              {{ line.origin }} →
              {{ line.destination }}
            </option>
          </select>

          <select
            v-model="newTrip.busId"
            required
          >
            <option
              value=""
              disabled
            >
              Selecione o ônibus
            </option>

            <option
              v-for="bus in buses"
              :key="bus.id"
              :value="bus.id"
            >
              {{ bus.prefix }} - {{ bus.model }}
            </option>
          </select>

          <input
            v-model="newTrip.tripDate"
            type="date"
            required
          />

          <input
            v-model="newTrip.tripTime"
            type="time"
            required
          />

          <input
            v-model.number="newTrip.passengers"
            type="number"
            min="0"
            placeholder="Passageiros"
            required
          />

          <button
            type="submit"
            :disabled="savingTrip"
          >
            {{ savingTrip ? 'Registrando...' : 'Registrar viagem' }}
          </button>
        </form>
      </section>

      <section>
        <h2>Viagens registradas</h2>

        <p v-if="trips.length === 0">
          Nenhuma viagem registrada.
        </p>

        <ul v-else>
          <li
            v-for="trip in trips"
            :key="trip.id"
          >
            <strong>
              {{ getLineDescription(trip.lineId) }}
            </strong>

            <br />

            Ônibus:
            {{ getBusDescription(trip.busId) }}

            <br />

            Data:
            {{ trip.tripDate }}

            |
            Horário:
            {{ trip.tripTime }}

            |
            Passageiros:
            {{ trip.passengers }}
          </li>
        </ul>
      </section>
    </template>
  </main>
</template>

<style scoped>
.container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 32px;
}

section {
  margin-bottom: 32px;
}

form {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

input,
select,
button {
  padding: 8px;
}

button {
  cursor: pointer;
}

button:disabled {
  cursor: not-allowed;
}

.error {
  color: #b00020;
}

ul {
  padding-left: 20px;
}

li {
  margin-bottom: 12px;
}
</style>

