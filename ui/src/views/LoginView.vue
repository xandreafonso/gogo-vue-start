<script setup lang="ts">
import { PhArrowCircleRight } from "@phosphor-icons/vue";
import { toTypedSchema } from '@vee-validate/zod';
import { useForm, useField } from 'vee-validate';
import * as z from 'zod';

const loginSchema = toTypedSchema(
  z.object({
    email: z.string({ required_error: "Informe o email" }).email('Deve ser um email válido'),
    password: z.string({ required_error: "Informe a senha" }).min(3, 'Informe a senha'),
  })
);

const { handleSubmit, errors } = useForm({
  validationSchema: loginSchema
});

const { value: email } = useField('email');
const { value: password } = useField('password');

const onSubmit = handleSubmit(values => {
  alert(JSON.stringify(values, null, 2));
})

</script>

<template>
  <div class="container">
    <div class="login">
      <h1 class="login__tag">Newsletter App</h1>

      <h2 class="login__title">Login</h2>

      <form @submit="onSubmit" class="login__form">
        <div class="login__field">
          <input type="text" name="email" placeholder="Seu email" v-model="email" />
  
          <span class="login__error">{{ errors.email }}</span>
        </div>
        
        <div class="login__field">
          <input type="password" name="password" placeholder="Senha" v-model="password"/>
  
          <span class="login__error">{{ errors.password }}</span>
        </div>
  
        <button type="submit">
          <PhArrowCircleRight />
          Entrar
        </button>
      </form>
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login {
  width: 90%;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;

  @include large {
    width: 600px;
  }  
  
  .login__tag {
    text-align: center;
    font-weight: var(--FF-SEMI);
    color: var(--BG-COLOR);
    background-color: var(--HIGHLIGHT-COLOR);
    
    padding: .2rem .8rem;
  }

  .login__title {
    font-size: 3rem;
    font-weight: var(--FF-SEMI);
  }

  .login__form {
    width: 100%;

    display: flex;
    flex-direction: column;
    gap: 1rem;

    .login__field {
      display: flex;
      flex-direction: column;

      .login__error {
        color: var(--DANGER)
      }
    }

    input {
      border: none;
      border-radius: var(--RADIUS);
      height: 3rem;
      padding-inline: .5rem;
      outline: none;

      &:hover, &:focus {
        outline: 2px solid var(--HIGHLIGHT-COLOR);
      }
    }

    button {
      border: none;
      border-radius: var(--RADIUS);

      font-size: 2rem;
      text-transform: uppercase;
      font-weight: var(--FF-SEMI);
      color: var(--COLOR);
      background-color: var(--MAIN-COLOR);
      
      cursor: pointer;

      padding: .5rem;
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 1rem;

      transition: transform .5s;

      &:hover {
        transform: scale(1.1);
      }
    }
  }
}
</style>


