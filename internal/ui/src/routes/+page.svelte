<script>
  import { onMount } from "svelte"

  let flags = []
  let newKey = ""
  let newValue = ""

  const fetchFlags = async () => {
    const response = await fetch("/api/v1/feature_flags")
    flags = await response.json()
  }

  const confirmDelete = (flag) => {
    const isConfirmed = confirm("Are you sure you want to delete the flag: " + flag + "?")
    if (isConfirmed) deleteFlag(flag)
  }


  const callApi = async (flag, method, body) => { 
	try{
		const response = await fetch(`/api/v1/feature_flags/${flag}`, {
			method: method,
			body: body 
		})
		if (!response.ok){
			throw new Error(response.status + ': ' + response.statusText)
		}
	}catch(error){
		console.log(error.message)
		const errorMessages = {
          'GET': 'Error fetching feature data',
          'POST': 'Error adding feature flag',
          'PUT': 'Error updating feature flag',
          'DELETE': 'Error deleting feature flag'
        }
        const errorMessage = errorMessages[method] || 'Error'
        alert(`${errorMessage}: ${error.message}`)
	}
  }

  const deleteFlag = async (flag) => {
	callApi(flag, 'DELETE', null)
	fetchFlags()
  }
  

  const resetItems = () => {
	newValue = ""
	newKey = ""
  }	

  const getFeatureData = async(flag) => {
     try {
       const response = await fetch(`/api/v1/feature_flags/${flag}`, {
         method: "GET"
       })
       if (!response.ok) {
         throw new Error("Failed to fetch feature data")
       } else {
	 	const responseBody = await response.text()
         newValue = responseBody
         newKey = flag
       }
     } catch (error) {
       console.error("Error fetching feature data:", error.message)
       alert("Error fetching feature data: " + error.message)
     }
  }

  const updateFlag = async () => {
	callApi(newKey, 'PUT', newValue, false)
	resetItems()	
	fetchFlags()
  }

  const saveFlag = async () => {
	callApi(newKey, 'POST', newValue, false)	
    resetItems()
    fetchFlags()
  }

  onMount(fetchFlags)
</script>

<h1>Feature Flags</h1>

<button
  class="btn btn-primary"
  data-bs-toggle="modal"
  data-bs-target="#createModal">Add Feature Flag</button
>

<ol class="list-group list-group-numbered">
  {#each flags as flag}
    <li
      class="list-group-item d-flex justify-content-between align-items-start"
    >
      <div class="ms-2 me-auto">
        <button
          class="btn"
          on:click={getFeatureData(flag)}
          data-bs-toggle="modal"
          data-bs-target="#editModal">{flag}</button>
      </div>

      <button class="btn btn-warning" on:click={() => confirmDelete(flag)}
        >delete</button>
    </li>
  {/each}
</ol>

<div
  class="modal fade"
  id="createModal"
  tabindex="-1"
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
			  bind:value={newKey}
		    />
		</div>
        <div class="input-group mb-3">
			<span class="input-group-text">value</span>
			<textarea 
			  type="text"
			  class="form-control"
			  placeholder="Enter feature flag value"
			  bind:value={newValue}
		    />
		</div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary"b on:click={resetItems} data-bs-dismiss="modal">Close</button>
        <button
          type="button"
          class="btn btn-primary"
          on:click={saveFlag}
          data-bs-dismiss="modal">Save changes</button>
      </div>
    </div>
  </div>
</div>

<div
  class="modal fade"
  id="editModal"
  tabindex="-1"
  aria-labelledby="editModalLabel"
  aria-hidden="true"
  data-bs-backdrop="static"
  data-bs-keyboard="false"
>
  <div class="modal-dialog modal-dialog-centered">
    <div class="modal-content">
      <div class="modal-header">
        <h1 class="modal-title fs-5" id="editModalLabel">Edit Feature Flag: {newKey}</h1>
      </div>
      <div class="modal-body">
        <div class="input-group mb-3">
			<span class="input-group-text">value</span>
			<textarea 
			  type="text"
			  class="form-control"
			  placeholder="Enter feature flag value"
			  bind:value={newValue}
		    />
		</div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" on:click={resetItems} data-bs-dismiss="modal">Close</button>
        <button
          type="button"
          class="btn btn-primary"
          on:click={updateFlag}
          data-bs-dismiss="modal">Save changes</button>
      </div>
    </div>
  </div>
</div>
