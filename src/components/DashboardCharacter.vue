<template>
  <div>
    <div class="box">
      <img :src="avatarUrl" />
      <div class="name">
        <h1 :class="charClass">{{ characterName }}</h1>
        <h2>&lt;{{guildName}}&gt;</h2>
      </div>

      <h3>{{level + " " + charClass}}</h3>
      <h4>{{itemLevel + " ILVL"}}</h4>
    </div>
  </div>
</template>

<script>
export default {
  name: "DashboardCharacter",
  props: {
    character: Object
  },
  computed : {
    characterName: function () {
      return this.character.name;
    },
    level: function() {
      return this.character.level;
    },
    avatarUrl: function() {
      return (
        "http://render-eu.worldofwarcraft.com/character/" +
        this.character.thumbnail
      );
    },
    itemLevel: function() {
      return this.character.items.averageItemLevelEquipped;
    },
    guildName: function() {
      return this.character.guild.name;
    },
    charClass: function() {
      return classToString(this.character.class);
    }
  }
};

function classToString(classID) { 
  var classes = ["Warrior", "Mage", "Paladin"];
  return classes[classID];
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
}

.box {
  position: relative;
  width: 600px;
  height: 200px;
  background-color: rgba(0, 0, 0, 0.8);
}

img {
  width: 128px;
  height: 128px;
  position: absolute;
  top: 36px;
  left: 32px;
}

.name {
  display: block;
  width: 300px;
  position: absolute;
  top: 36px;
  left: 192px;
}

h1 {
  display: block;
  position: absolute;
  font-size: 2rem;
  top: 0;
  left: 0;
}

h2 {
  color: #c7c7c7;
  font-size: 1.25rem;
  margin-top: 10px;
  margin-left: 32px;
}

h3 {
  position: absolute;
  top: 72px;
  left: 192px;
  color: #d6ab1b;
  font-size: 1rem;
}

h4 {
  position: absolute;
  top: 48px;
  right: 16px;
  color: #d6ab1b;
  font-size: 1.25rem;
}

.Paladin {
  color: #f58cba;
}
</style>
