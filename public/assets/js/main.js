// import App from "../../../views/App.vue";


const app = Vue.createApp({
    // render: h => h(App),
    // delimiters: ['${', '}']
}).mount('#storage')

console.log(app);

// Определяем новый глобальный компонент с именем button-counter
// storage.component('App', {
//     data() {
//         return {
//             count: 0
//         }
//     },
//     template: `
//       <button @click="count++">
//       Счётчик кликов — {{ count }}
//       </button>`
// })