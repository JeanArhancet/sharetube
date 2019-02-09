<template>
  <v-form>
    <v-container>
      <v-layout>
        <v-flex xs12 sm6>
          <v-text-field label="Regular" single-line></v-text-field>
        </v-flex>
        <v-flex xs12 sm6>
          <v-text-field
            v-model="password"
            :append-icon="show1 ? 'visibility_off' : 'visibility'"
            :rules="[rules.required, rules.min]"
            :type="show1 ? 'text' : 'password'"
            name="input-10-1"
            label="Normal with hint text"
            hint="At least 8 characters"
            counter
            @click:append="show1 = !show1"
          ></v-text-field>
        </v-flex>
      </v-layout>
    </v-container>
  </v-form>
</template>
