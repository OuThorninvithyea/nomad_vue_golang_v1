import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useCities() {
    const cities = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/cities')
            cities.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { cities, loading, error }
}
