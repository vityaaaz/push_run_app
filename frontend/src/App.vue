<template>
  <ion-app>
    <ion-router-outlet />
  </ion-app>
</template>

<script setup lang="ts">
import { IonApp, IonRouterOutlet } from '@ionic/vue'
</script>

<style>
ion-app {
  background-color: var(--ion-background-color);
}
</style>
