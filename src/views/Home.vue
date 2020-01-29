<template>
  <div class="home">
    <CharShortInfo />
  </div>
</template>

<script>
// @ is an alias to /src
import CharShortInfo from "@/components/CharShortInfo.vue";
export default {
  name: "home",
  components: {
    CharShortInfo
  }
};
</script>

<style>
body {
  background-color: #5d5d5d;
}
</style>
