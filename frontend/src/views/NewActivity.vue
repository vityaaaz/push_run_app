<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-buttons slot="start">
          <ion-back-button default-href="/home"></ion-back-button>
        </ion-buttons>
        <ion-title>Новая активность</ion-title>
      </ion-toolbar>
    </ion-header>

    <ion-content class="ion-padding">
      <ion-grid>
        <ion-row>
          <ion-col size="12">
            <ion-segment v-model="activityType">
              <ion-segment-button value="run">
                <ion-label>Бег</ion-label>
              </ion-segment-button>
              <ion-segment-button value="bike">
                <ion-label>Велосипед</ion-label>
              </ion-segment-button>
            </ion-segment>
          </ion-col>
        </ion-row>

        <ion-row>
          <ion-col size="12">
            <ion-card>
              <ion-card-content>
                <div class="activity-stats">
                  <div class="stat">
                    <h3>{{ formatTime(elapsedTime) }}</h3>
                    <p>Время</p>
                  </div>
                  <div class="stat">
                    <h3>{{ distance.toFixed(2) }} км</h3>
                    <p>Дистанция</p>
                  </div>
                  <div class="stat">
                    <h3>{{ pace.toFixed(2) }} мин/км</h3>
                    <p>Темп</p>
                  </div>
                </div>
              </ion-card-content>
            </ion-card>
          </ion-col>
        </ion-row>

        <ion-row>
          <ion-col size="12">
            <ion-button
              expand="block"
              :color="isTracking ? 'danger' : 'success'"
              @click="toggleTracking"
            >
              {{ isTracking ? 'Остановить' : 'Начать' }}
            </ion-button>
          </ion-col>
        </ion-row>
      </ion-grid>
    </ion-content>
  </ion-page>
</template>

<script setup lang="ts">
import { ref, computed, onUnmounted } from 'vue'
import {
  IonPage,
  IonHeader,
  IonToolbar,
  IonTitle,
  IonContent,
  IonGrid,
  IonRow,
  IonCol,
  IonCard,
  IonCardContent,
  IonButton,
  IonSegment,
  IonSegmentButton,
  IonLabel,
  IonBackButton,
  IonButtons
} from '@ionic/vue'

const activityType = ref('run')
const isTracking = ref(false)
const startTime = ref<Date | null>(null)
const elapsedTime = ref(0)
const distance = ref(0)
const timer = ref<number | null>(null)

const pace = computed(() => {
  if (distance.value === 0) return 0
  return (elapsedTime.value / 60) / distance.value
})

const formatTime = (seconds: number) => {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60
  return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
}

const toggleTracking = () => {
  isTracking.value = !isTracking.value
  if (isTracking.value) {
    startTime.value = new Date()
    timer.value = window.setInterval(() => {
      if (startTime.value) {
        elapsedTime.value = Math.floor((new Date().getTime() - startTime.value.getTime()) / 1000)
        // Simulate distance increase
        distance.value += 0.01
      }
    }, 1000)
  } else {
    if (timer.value) {
      clearInterval(timer.value)
      timer.value = null
    }
  }
}

onUnmounted(() => {
  if (timer.value) {
    clearInterval(timer.value)
  }
})
</script>

<style scoped>
.activity-stats {
  display: flex;
  justify-content: space-around;
  text-align: center;
  margin: 20px 0;
}

.stat h3 {
  margin: 0;
  font-size: 1.5em;
}

.stat p {
  margin: 5px 0 0;
  color: var(--ion-color-medium);
  font-size: 0.9em;
}
</style> 