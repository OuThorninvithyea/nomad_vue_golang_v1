import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useLogoDropdown() {
    const logoDropdown = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/logo-dropdown')
            logoDropdown.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { logoDropdown, loading, error }
}
