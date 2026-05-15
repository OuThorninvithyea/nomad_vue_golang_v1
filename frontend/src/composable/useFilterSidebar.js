import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useFilterSidebar() {
    const filterSidebar = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/filter-sidebar')
            filterSidebar.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { filterSidebar, loading, error }
}
