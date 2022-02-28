<template>
  <div class="form-group">
      <label :for="Id">{{labelName}}</label>                                               
      <input  :class="ClassValue" :id="Id" :name="Id" :type="Inputtype" :placeholder="PlaceHolder" :disabled="Disable">
  </div>
</template>
<script>
export default {
  name: "InputBase",
  data() {
    return {
    };
  },
  props: {
    Id:{
      type: String,
      default: "text",
      description: "",
    },
    ClassValue:{
      type: String,
      default: "form-control",
      description: "",
    },
    labelName:{
      type: String,
      default: "",
      description: "",
    },
    Inputtype: {
      type: String,
      default: "text",
      description: "",
    },
    Disable: {
      type: Boolean,
      default: false,
      description: "",
    },
    PlaceHolder:{
      type: String,
      default: "",
      description: "",
    },
  },
  created(){
    console.log("PROPS INPUT",this.$data,this.labelName)
  },
  watch:{
    
  },
  methods: {
  },
};
</script>
<style scoped>

</style>