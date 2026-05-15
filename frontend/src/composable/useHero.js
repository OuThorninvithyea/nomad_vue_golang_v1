import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useHero() {
    const hero = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/hero')
            hero.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { hero, loading, error }
}
