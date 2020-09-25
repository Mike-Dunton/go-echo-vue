<template>
  <b-container>
      <b-row>
          <b-col md=6>
              <h2>My Tasks -  <b-button pill size="sm" variant="outline-primary" v-on:click='isTasksReversed = !isTasksReversed'><b-icon icon="arrow-down-up"></b-icon></b-button></h2>
              <b-list-group class="list-group">
                  <b-list-group-item class="d-flex justify-content-between align-items-center" v-for="(task, index) in cTasks" :key="task.id">
                        {{ task.description }}                       
                            <b-icon v-on:click="deleteTask(index)" icon="trash2" variant="danger" ></b-icon>
                  </b-list-group-item>
              </b-list-group>
              <div>
                <b-input-group prepend="New Task" class="mt-3">
                    <b-form-input type="text" 
                        v-on:keyup.enter="createTask"
                        v-model="newTask.description"></b-form-input>
                    <b-input-group-append>
                        <b-button variant="outline-success" v-on:click="createTask">Create</b-button>
                    </b-input-group-append>
                </b-input-group>
              </div><!-- /input-group -->
          </b-col>
      </b-row>
  </b-container>
</template>

<script>
import $ from 'jquery'
import axios from 'axios'

const transport = axios.create({
  withCredentials: true
})

export default {
  name: 'Todo',
  data: function() {
    return {
      isTasksReversed: false,
      tasks: [],
      newTask: {}
    }
  },
  // This is run whenever the page is loaded to make sure we have a current task list
  created: function() {
      transport.get('/auth/tasks').then((response) => {
          this.tasks = response.data.items ? response.data.items : []
      })
  },
  computed: {
      cTasks: function() {
          if(this.isTasksReversed) {
              return this.tasks.slice().reverse();
          } else {
              return this.tasks
          }
          
      }     
  },
  methods: {
      createTask: function() {
          if (!$.trim(this.newTask.description)) {
              this.newTask = {}
              return
          }

      // Post the new task to the /tasks route using the $http client
         transport.put('/auth/tasks', this.newTask).then((response) => {
              this.newTask.id = response.created
              this.tasks.push(this.newTask)
              console.log("Task created!")
              console.log(this.newTask)
              this.newTask = {}
          }).catch((error) => {
              console.log(error)
          });
      },

      deleteTask: function(index) {
// Use the $http client to delete a task by its id
          transport.delete('/auth/tasks/' + this.cTasks[index].id).then(() => {
              transport.get('/auth/tasks').then((response) => {
                  this.tasks = response.data.items ? response.data.items : []
              }).catch((error) => {console.log(error)})
          }).catch((error) => {
              console.log(error)
          })
      }                        
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
b-list-group {
  list-style-type: none;
  padding: 0;
}
b-list-group-item {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
