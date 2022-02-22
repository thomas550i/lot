<template>
<div class="form-group pd-none">
  <label :for="Id" class="col-sm-3 control-label text-darkness">{{labelName}}</label>                                               
  <div class="col-sm-8">
      <input 
        :id="Id"
        :type="Inputtype"
        :disabled="Disable"
        :autofocus="AutoFocus"
        :value="Value"
      >
      <span :id="Id+labelName" v-if="CheckError" class="">
        {{ErrorMsg}}
      </span>
  </div>
  
</div>
</template>
<script>
export default {
  name: "Input",
  data() {
    return {
      CheckError:false,
      visible: true,
      ErrorMsg:"",
      type:"",
    };
  },
  props: {
    Id:{
      type: String,
      default: "text",
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
    AutoFocus: {
      type: Boolean,
      default: false,
      description: "",
    },
    Value:{
      type: [String,Number,Boolean,Object],
      description: "Override Specific value",
    },
    MaxLength:{
      type: [Number,Date],
      description: "Check Maximum Length of Given Value",
    },
    MinLength:{
      type: [Number,Date],
      description: "Check Minimum Length of Given Value",
    },
    Validate:{
      type: Boolean,
      default: false,
      description: "Validation if required",
    }
  },
  
  beforeCreate(){
  },
  watch:{
    Validate(Value,oldValue){
      if(Value===true){
        switch(this.$data.Value.type){
          case String:
            if(!(this.$data.Value<this.$data.MaxLength)){
               this.CheckError = true
               this.ErrorMsg = "Field "+this.$data.labelName+"Cannot Be more than "+this.$data.MaxLength
            }
            if(!this.$data.Value>this.$data.MinLength){
              this.CheckError = true
              this.ErrorMsg = "Field "+this.$data.labelName+"Cannot Be less than "+this.$data.MaxLength;
            }
            break;
          case Number:
            if(!(this.$data.Value<this.$data.MaxLength)){
               this.CheckError = true
               this.ErrorMsg = "Field "+this.$data.labelName+"Cannot Be more than "+this.$data.MaxLength
            }
            if(!this.$data.Value>this.$data.MinLength){
              this.CheckError = true
              this.ErrorMsg = "Field "+this.$data.labelName+"Cannot Be less than "+this.$data.MaxLength
            }
            break;

          case Date:
            if(!(this.$data.Value<this.$data.MaxLength)){
               this.CheckError = true
               this.ErrorMsg = "Field "+this.$data.labelName+"Cannot Be more than "+this.$data.MaxLength
            }
            if(!this.$data.Value>this.$data.MinLength){
              this.CheckError = true
              this.ErrorMsg = "Field "+this.$data.labelName+"Cannot Be less than "+this.$data.MaxLength
            }
            break;

        }

      }else{
        console.log("OLD VALUE",oldValue)
      } 
    },
  },
  methods: {
  },
};
</script>
<style scoped>

</style>