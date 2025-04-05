<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>ПУШ</ion-title>
      </ion-toolbar>
    </ion-header>

    <ion-content class="ion-padding">
      <ion-grid>
        <ion-row>
          <ion-col size="12">
            <ion-card>
              <ion-card-header>
                <ion-card-title>Добро пожаловать в ПУШ</ion-card-title>
                <ion-card-subtitle>Трекер для бега и велоспорта</ion-card-subtitle>
              </ion-card-header>

              <ion-card-content>
                <ion-button expand="block" router-link="/activities/new">
                  Начать тренировку
                </ion-button>
              </ion-card-content>
            </ion-card>
          </ion-col>
        </ion-row>

        <ion-row>
          <ion-col size="12">
            <ion-list>
              <ion-list-header>
                <ion-label>Последние активности</ion-label>
              </ion-list-header>
              <ion-item v-for="activity in recentActivities" :key="activity.id" router-link="/activities">
                <ion-label>
                  <h2>{{ activity.type }}</h2>
                  <p>{{ activity.date }}</p>
                  <p>{{ activity.distance }} км</p>
                </ion-label>
              </ion-item>
            </ion-list>
          </ion-col>
        </ion-row>
      </ion-grid>
    </ion-content>

    <ion-footer>
      <ion-tabs>
        <ion-tab-bar slot="bottom">
          <ion-tab-button tab="home" selected>
            <ion-icon :icon="home"></ion-icon>
            <ion-label>Главная</ion-label>
          </ion-tab-button>
          <ion-tab-button tab="activities">
            <ion-icon :icon="list"></ion-icon>
            <ion-label>Активности</ion-label>
          </ion-tab-button>
          <ion-tab-button tab="profile">
            <ion-icon :icon="person"></ion-icon>
            <ion-label>Профиль</ion-label>
          </ion-tab-button>
        </ion-tab-bar>
      </ion-tabs>
    </ion-footer>
  </ion-page>
</template>

<script setup lang="ts">
import { ref } from 'vue'
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
  IonCardHeader,
  IonCardTitle,
  IonCardSubtitle,
  IonCardContent,
  IonButton,
  IonList,
  IonListHeader,
  IonItem,
  IonLabel,
  IonFooter,
  IonTabs,
  IonTabBar,
  IonTabButton,
  IonIcon
} from '@ionic/vue'
import { home, list, person } from 'ionicons/icons'

const recentActivities = ref([
  {
    id: 1,
    type: 'Бег',
    date: '2024-04-04',
    distance: 5.2
  },
  {
    id: 2,
    type: 'Велосипед',
    date: '2024-04-03',
    distance: 15.7
  }
])
</script>

<style scoped>
ion-card {
  margin-bottom: 20px;
}

ion-list {
  margin-top: 20px;
}
</style> 