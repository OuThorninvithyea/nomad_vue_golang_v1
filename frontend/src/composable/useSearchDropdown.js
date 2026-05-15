import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useSearchDropdown() {
    const searchDropdown = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/search-dropdown')
            searchDropdown.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { searchDropdown, loading, error }
}
