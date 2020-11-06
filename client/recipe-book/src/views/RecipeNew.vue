<template>
  <div class="recipe-single">
    <section class="hero is-warning">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            New Recipe
          </h1>
        </div>
      </div>
    </section>
    <section>
      <div class="container">
        <div class="container">
          <div class="notification is-primary">
            <p v-if="errors.length">
              <b>Please correct the following error(s):</b>
              <ul>
                <li v-for="error in errors" :key="error">{{ error }}</li>
              </ul>
            </p>
          </div>
        </div>
          

          <!--Name Field-->
          <div class="field is-horizontal">
            <div class="field-label is-normal">
              <label class="label">Name</label>
            </div>
            <div class="field-body">
              <div class="control is-expanded">
                <input class="input" type="text" v-model.trim="name" placeholder="Fried Jalapeno Poppers">
              </div>
            </div>
          </div>

          <div class="field is-horizontal">
            <div class="field-label is-normal">
              <label class="label">Category</label>
            </div>
            <div class="field-body">
              <div class="control">
                <div class="select">
                  <select v-model="category">
                    <option>Food</option>
                    <option>Drink</option>
                  </select>
                </div>
              </div>
            </div>
          </div>

          <div class="field is-horizontal">
            <div class="field-label is-normal">
              <label class="label">Notes</label>
            </div>
            <div class="field-body">
              <div class="field">
                <div class="control">
                  <textarea v-model="notes" class="textarea" placeholder="Notes/Procedure"></textarea>
                </div>
              </div>
            </div>
          </div>

          <!--UPLOAD-->
          <form enctype="multipart/form-data" novalidate v-if="isInitial || isSaving">
            <div class="field is-horizontal">
              <div class="field-label is-normal">
                <label class="label">Upload images</label>
              </div>
              <div class="field-body">
                <div class="field">
                  <div class="dropbox">
                    <input type="file" multiple :name="uploadFieldName" :disabled="isSaving" @change="filesChange($event.target.name, $event.target.files); fileCount = $event.target.files.length"
                      accept="image/*" class="input-file">
                      <p v-if="isInitial">
                        Drag your file(s) here to begin<br> or click to browse
                      </p>
                      <p v-if="isSaving">
                        Uploading {{ fileCount }} files...
                      </p>
                  </div>
                </div>
              </div>
            </div>
          </form>

          <!--SUCCESS-->
          <div v-if="isSuccess">
            <h2>Uploaded {{ previewedFiles.length }} file(s) successfully.</h2>
            <p>
              <a href="javascript:void(0)" @click="reset()">Upload again</a>
            </p>
            <div class="columns is-multiline is-mobile">
              <div v-for="item in previewedFiles" :key="item.id" class="column is-one-quarter">
                <figure class="image is-square">
                  <img :src="item.url"  :alt="item.originalName">
                </figure>
              </div>
            </div>
          </div>
          <!--FAILED-->
          <div v-if="isFailed">
            <h2>Uploaded failed.</h2>
            <p>
              <a href="javascript:void(0)" @click="reset()">Try again</a>
            </p>
            <pre>{{ uploadError }}</pre>
          </div>

          <div class="field is-horizontal is-grouped mt-2">
            <div class="field-label">
              <!-- Left empty for spacing -->
            </div>
            <div class="field-body">
              <div class="control">
                <button v-on:click="validateRecipe()" class="button is-link">Submit</button>
              </div>
              <div class="control">
                <button  v-on:click="reset()" class="button is-link is-light">Cancel</button>
              </div>
            </div>
          </div>
      </div>
    </section>
  </div>
</template>
<script>
  import { previewUpload } from '../services/file-upload.preview.service.js'; // preview service
  import { recipeService } from '../services/recipes.js';   // real service

  const STATUS_INITIAL = 0, STATUS_SAVING = 1, STATUS_SUCCESS = 2, STATUS_FAILED = 3;
  export default {
    name: 'RecipeNew',
     data() {
      return {
        name: "",
        category: "",
        notes: "",
        previewedFiles: [],
        uploadError: null,
        currentStatus: null,
        uploadFieldName: 'photos',
        errors: [],
        files: []
      }
    },
    computed: {
      isInitial() {
        return this.currentStatus === STATUS_INITIAL;
      },
      isSaving() {
        return this.currentStatus === STATUS_SAVING;
      },
      isSuccess() {
        return this.currentStatus === STATUS_SUCCESS;
      },
      isFailed() {
        return this.currentStatus === STATUS_FAILED;
      }
    },  
methods: {
      reset() {
        // reset form to initial state
        this.currentStatus = STATUS_INITIAL;
        this.previewedFiles = [];
        this.files = [];
        this.name = "";
        this.notes = "";
        this.errors = [];
        this.category = "";
        this.uploadError = null;
      },
      previewImages(formData) {
        previewUpload(formData)
          .then((x) => {
                return new Promise(resolve => setTimeout(() => resolve(x), 1500));
          })
          .then(x => {
            this.previewedFiles = [].concat(x);
            this.currentStatus = STATUS_SUCCESS;
          })
          .catch(err => {
            this.uploadError = err.response;
            this.currentStatus = STATUS_FAILED;
          });
      },
      validateRecipe() {
        this.errors = [];
        if (this.name == "") {
          this.errors.push("Recipe must have a name")
        }
        if (this.category == "") {
          this.errors.push("You must select a category")
        }
        
        if (this.errors.length == 0) {
          this.save()
        }
      },
      save() {
        // upload data to the server
        this.currentStatus = STATUS_SAVING;
        const formData = new FormData();
        if (this.files.length > 0) {
        Array
          .from(Array(this.files.length).keys())
          .map(x => {
            formData.append("images", this.files[x], this.files[x].name);
          });
        }
        formData.append("name", this.name)
        formData.append("category", this.category)
        formData.append("notes", this.notes)
        recipeService.postRecipe(formData, response => {
            if(response.length != 0) {
              this.currentStatus = STATUS_SUCCESS;
            } else {
              this.currentStatus = STATUS_FAILED;
            }  
        })
      },
      filesChange(fieldName, fileList) {
        // handle file changes
        const formData = new FormData();
        if (!fileList.length) return;
        // append the files to FormData
        this.files = fileList
        Array
          .from(Array(fileList.length).keys())
          .map(x => {
            formData.append(fieldName, fileList[x], fileList[x].name);
          });
        // save it
        this.previewImages(formData);
      }
    },
    mounted() {
      this.reset();
    },
  }
</script>
<style lang="scss" scoped>
  .hero {
    margin-bottom: 25px;
  }

  .control.is-expanded {
    flex-grow: 1;
    flex-shrink: 1;
  }

 .dropbox {
    outline: 2px dashed grey; /* the dash box */
    outline-offset: -10px;
    background: lightcyan;
    color: dimgray;
    padding: 10px 10px;
    min-height: 200px; /* minimum height */
    position: relative;
    cursor: pointer;
  }
  
  .input-file {
    opacity: 0; /* invisible but it's there! */
    width: 100%;
    height: 200px;
    position: absolute;
    cursor: pointer;
  }
  
  .dropbox:hover {
    background: lightblue; /* when mouse over to the drop zone, change color */
  }
  
  .dropbox p {
    font-size: 1.2em;
    text-align: center;
    padding: 50px 0;
  }
</style>