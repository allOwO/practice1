import Vue from 'vue'
import VueRouter from 'vue-router'
import Add from '../components/Add'
import Change from '../components/Change'

Vue.use(VueRouter)

const routes = [
  {
    path: '/add',
    component: Add
  }, {
    path: '/change',
    component: Change
  }
]
const router = new VueRouter({
  routes
})
export default router
