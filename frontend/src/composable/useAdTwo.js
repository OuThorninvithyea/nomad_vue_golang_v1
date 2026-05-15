import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useAdTwo() {
    const adTwo = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/ad-two')
            adTwo.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { adTwo, loading, error }
}
