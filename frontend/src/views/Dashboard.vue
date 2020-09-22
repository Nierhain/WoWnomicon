<template>
<v-row>
  <DashboardCharacter v-for="character in characters" :key="character" :character="character"/>
</v-row>
</template>

<script>
// @ is an alias to /src
import DashboardCharacter from '@/components/DashboardCharacter.vue'

export default {
  name: 'Dashboard',
  components: {
    DashboardCharacter
  },
  data() {
    return {
      characters: Array
    };
  },
  mounted() {
      window.backend.getDummy().then(result => {
        this.characters = JSON.parse(result);
      })

      
  }
}
</script>
