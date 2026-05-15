import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useSuggest() {
    const suggest = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/suggest')
            suggest.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { suggest, loading, error }
}
