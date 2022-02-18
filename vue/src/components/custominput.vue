<template>
    <input v-if="value!=0" type="text" id="cinput" @keypress="isNumber($event)" class="cutomeinput"  :value="value" @input="onInput" />
    <input v-if="value==0" type="text"  @keypress="isNumber($event)" class="cutomeinput" @input="onInput" />
</template>

<script>
export default {
  name: 'CustomInput',
  props: {
    value: {
      type: Number,
      required: true,
    },
  },
  created(){
        if( this.value == 0 ){
           console.log("")
    }
  },
  computed: {
  },
  methods: {
       onInput(event) {
           console.log("event",event.target.value)
           if (event.target.value==""){
                this.$emit('updated', "");
           } else {
               this.$emit('updated', parseInt(event.target.value));
           }
      
    },
       isNumber (evt){
        const keysAllowed= ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9'];
        const keyPressed = evt.key;
        
        if (!keysAllowed.includes(keyPressed)) {
              evt.preventDefault()
        }
      },
  },
}
</script>