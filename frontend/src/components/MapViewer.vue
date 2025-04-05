<template>
  <div class="map-container">
    <div ref="mapContainer" class="map"></div>
    <div class="controls">
      <button @click="startTracking" :disabled="isTracking">
        Начать трекинг
      </button>
      <button @click="stopTracking" :disabled="!isTracking">
        Остановить трекинг
      </button>
      <button @click="clearTrack">
        Очистить маршрут
      </button>
    </div>
    <div v-if="weather" class="weather-info">
      <h3>Погода</h3>
      <p>Температура: {{ weather.temperature }}°C</p>
      <p>Ощущается как: {{ weather.feelsLike }}°C</p>
      <p>Условия: {{ weather.condition }}</p>
      <p>Ветер: {{ weather.windSpeed }} м/с</p>
      <p>Влажность: {{ weather.humidity }}%</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { mapService, MapSettings, Weather } from '@/services/mapService'

const mapContainer = ref<HTMLElement | null>(null)
const isTracking = ref(false)
const weather = ref<Weather | null>(null)

const settings: MapSettings = {
  center: {
    latitude: 55.751244,
    longitude: 37.618423
  },
  zoom: 12,
  mapType: 'map',
  showTraffic: true
}

onMounted(async () => {
  if (mapContainer.value) {
    await mapService.initMap(mapContainer.value.id, settings)
  }
})

onUnmounted(() => {
  mapService.stopTracking()
})

const startTracking = async () => {
  isTracking.value = true
  mapService.startTracking()
  
  // Получаем погоду для текущего местоположения
  const position = await new Promise<GeolocationPosition>((resolve, reject) => {
    navigator.geolocation.getCurrentPosition(resolve, reject)
  })
  
  weather.value = await mapService.getWeather({
    latitude: position.coords.latitude,
    longitude: position.coords.longitude
  })
}

const stopTracking = () => {
  isTracking.value = false
  mapService.stopTracking()
}

const clearTrack = () => {
  mapService.clearTrack()
}
</script>

<style scoped>
.map-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.map {
  width: 100%;
  height: 100%;
}

.controls {
  position: absolute;
  top: 10px;
  left: 10px;
  z-index: 1000;
  display: flex;
  gap: 10px;
}

.controls button {
  padding: 8px 16px;
  background-color: #fff;
  border: 1px solid #ccc;
  border-radius: 4px;
  cursor: pointer;
}

.controls button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.weather-info {
  position: absolute;
  top: 10px;
  right: 10px;
  z-index: 1000;
  background-color: #fff;
  padding: 10px;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.weather-info h3 {
  margin: 0 0 10px 0;
}

.weather-info p {
  margin: 5px 0;
}
</style> 