import { useEffect, useState } from 'preact/hooks'
import { FunctionalComponent } from 'preact'
import './style.sass'

type Flag = string

const Page: FunctionalComponent = () => {
  const [flags, setFlags] = useState<Flag[]>([])
  const [newKey, setNewKey] = useState('')
  const [newValue, setNewValue] = useState('')
  const path = window.location.origin

  const fetchFlags = async () => {
    try {
      const response = await fetch('/api/v1/feature_flags')
      const data = await response.json()
      setFlags(data ?? [])
    } catch (error) {
      alert('Error fetching feature flags')
    }
  }

  const confirmDelete = (flag: string) => {
    if (window.confirm(`Are you sure you want to delete the flag: ${flag}?`)) {
      deleteFlag(flag)
    }
  }

  const callApi = async (flag: string, method: string, body: string | null) => {
    try {
      const response = await fetch(`/api/v1/feature_flags/${flag}`, {
        method,
        body,
      })
      if (!response.ok) {
        throw new Error(response.status + ': ' + response.statusText)
      }
    } catch (error: any) {
      const errorMessages: Record<string, string> = {
        GET: 'Error fetching feature data',
        POST: 'Error adding feature flag',
        PUT: 'Error updating feature flag',
        DELETE: 'Error deleting feature flag',
      }
      const errorMessage = errorMessages[method] || 'Error'
      alert(`${errorMessage}: ${error.message}`)
    }
  }

  const deleteFlag = async (flag: string) => {
    await callApi(flag, 'DELETE', null)
    fetchFlags()
  }

  const resetItems = () => {
    setNewKey('')
    setNewValue('')
  }

  const getFeatureData = async (flag: string) => {
    try {
      const response = await fetch(`/api/v1/feature_flags/${flag}`, {
        method: 'GET',
      })
      if (!response.ok) {
        throw new Error('Failed to fetch feature data')
      } else {
        const responseBody = await response.text()
        setNewValue(responseBody)
        setNewKey(flag)
      }
    } catch (error: any) {
      alert('Error fetching feature data: ' + error.message)
    }
  }

  const updateFlag = async () => {
    await callApi(newKey, 'PUT', newValue)
    resetItems()
    fetchFlags()
  }

  const saveFlag = async () => {
    await callApi(newKey, 'POST', newValue)
    resetItems()
    fetchFlags()
  }

  useEffect(() => {
    fetchFlags()
  }, [])

  // Bootstrap modals require manual show/hide if not using data-bs-toggle
  // We'll keep the data-bs-toggle/data-bs-target attributes for compatibility

  return (
    
    <div class="container py-4">
      <button
        class="btn btn-primary mb-5"
        data-bs-toggle="modal"
        data-bs-target="#createModal"
      >
        Add Feature Flag
      </button>

      <ul class="list-group">
        {flags.map((flag) => (
          <li class="list-group-item d-flex justify-content-between align-items-start" key={flag}>
            <div class="ms-2 me-auto">
              <div class="fw-bold">
                <button
                  class="btn btn-outline-info"
                  data-bs-toggle="modal"
                  data-bs-target="#editModal"
                  onClick={() => getFeatureData(flag)}
                >
                  {flag}
                </button>
              </div>
              <i>
                curl -X 'GET' '{path}/api/v1/feature_flags/{flag}'
              </i>
            </div>
            <button class="btn btn-danger" onClick={() => confirmDelete(flag)}>
              delete
            </button>
          </li>
        ))}
      </ul>

      {/* Create Modal */}
      <div
        class="modal fade"
        id="createModal"
        tabIndex={-1}
        aria-labelledby="createModalLabel"
        aria-hidden="true"
        data-bs-backdrop="static"
        data-bs-keyboard="false"
      >
        <div class="modal-dialog modal-dialog-centered">
          <div class="modal-content">
            <div class="modal-header">
              <h1 class="modal-title fs-5" id="createModalLabel">
                Create Feature Flag
              </h1>
            </div>
            <div class="modal-body">
              <div class="input-group mb-3">
                <span class="input-group-text">key</span>
                <input
                  type="text"
                  class="form-control"
                  placeholder="Enter feature flag key"
                  value={newKey}
                  onInput={(e) => setNewKey((e.target as HTMLInputElement).value)}
                />
              </div>
              <div class="input-group mb-3">
                <span class="input-group-text">value</span>
                <textarea
                  class="form-control"
                  placeholder="Enter feature flag value"
                  value={newValue}
                  onInput={(e) => setNewValue((e.target as HTMLTextAreaElement).value)}
                />
              </div>
            </div>
            <div class="modal-footer">
              <button
                type="button"
                class="btn btn-secondary"
                onClick={resetItems}
                data-bs-dismiss="modal"
              >
                Close
              </button>
              <button
                type="button"
                class="btn btn-primary"
                onClick={saveFlag}
                data-bs-dismiss="modal"
              >
                Save changes
              </button>
            </div>
          </div>
        </div>
      </div>

      {/* Edit Modal */}
      <div
        class="modal fade"
        id="editModal"
        tabIndex={-1}
        aria-labelledby="editModalLabel"
        aria-hidden="true"
        data-bs-backdrop="static"
        data-bs-keyboard="false"
      >
        <div class="modal-dialog modal-dialog-centered">
          <div class="modal-content">
            <div class="modal-header">
              <h1 class="modal-title fs-5" id="editModalLabel">
                Edit Feature Flag: {newKey}
              </h1>
            </div>
            <div class="modal-body">
              <div class="input-group mb-3">
                <span class="input-group-text">value</span>
                <textarea
                  class="form-control"
                  placeholder="Enter feature flag value"
                  value={newValue}
                  onInput={(e) => setNewValue((e.target as HTMLTextAreaElement).value)}
                />
              </div>
            </div>
            <div class="modal-footer">
              <button
                type="button"
                class="btn btn-secondary"
                onClick={resetItems}
                data-bs-dismiss="modal"
              >
                Close
              </button>
              <button
                type="button"
                class="btn btn-primary"
                onClick={updateFlag}
                data-bs-dismiss="modal"
              >
                Save changes
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    
  )
}

export default Page