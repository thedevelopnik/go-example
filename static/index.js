Vue.use(Vuetify)

new Vue({
  el: '#app',
  data() {
    return {
      newTodo: '',
      todos: []
    }
  },
  methods: {
    submit: function() {
      this.todos.push(this.newTodo)
      this.newTodo = ''
    },
    remove: function(index) {
      this.todos.splice(index, (index + 1))
    },
    getTodos: function() {
      axios.get('/gettodos')
      .then(res => {
        res.data.forEach(todo => {
          this.todos.push(todo.name)
        })
      })
    }
  }
})
