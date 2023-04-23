<template>
  <v-card id="app-layout">
    <v-layout>
      <v-app-bar
        title="Application bar"
        dense
        :style="{ borderBottom: '1px solid #eee' }"
      ></v-app-bar>

      <v-navigation-drawer :style="{ border: '1px solid #eee', padding: 0 }">
        <v-list
          :style="{ padding: 0, textAlign: 'center' }"
          :density="'compact'"
        >
          <router-link to="/">
            <v-list-item title="Home" :style="{ textAlign: 'left' }" />
          </router-link>

          <router-link to="/about">
            <v-list-item title="About" :style="{ textAlign: 'left' }" />
          </router-link>
        </v-list>
      </v-navigation-drawer>

      <v-main>
        <div :style="{ height: '100%', padding: '4px 16px' }">
          <router-view />
        </div>
      </v-main>
    </v-layout>
  </v-card>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "AppLayout",
});
</script>
