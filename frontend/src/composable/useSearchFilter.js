import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useSearchFilter() {
    const searchFilter = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/search-filter')
            searchFilter.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { searchFilter, loading, error }
}
