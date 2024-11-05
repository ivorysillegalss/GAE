<template>
  <div class="back">
    <div class="logo-big">
      <img src="/logo.svg" class="logo" alt="Vite logo" />
    </div>
    <div class="buttons">
      <button>首页</button>
      <button>排名</button>
      <button>开发者Nation</button>
      <button>评估报告</button>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.back {
  /* width: 100%; */
  height: 150px;
  display: flex;
  position: absolute;
  top: 0px;
  /* flex-direction: row; */
  align-items: center;
}
.logo-big {
  cursor: pointer;
  position: relative;
  right: 460px;
}
.logo {
  height: 3em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #ffa61f);
}
.buttons > button{
  margin-left: 10px;
}
</style>
