<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Customer Management</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
    <div class="container my-5">
        <h1 class="mb-4">Customer Management</h1>

        @if (session('success'))
            <div class="alert alert-success">
                {{ session('success') }}
            </div>
        @endif

        @if ($errors->any())
            <div class="alert alert-danger">
                <ul>
                    @foreach ($errors->all() as $error)
                        <li>{{ $error }}</li>
                    @endforeach
                </ul>
            </div>
        @endif

        <div class="card mb-5">
            <div class="card-header">
                <h3>Add Customer</h3>
            </div>
            <div class="card-body">
                <form id="createCustomerForm" action="/" method="POST">
                    @csrf
                    <div class="row">
                        <div class="col-md-6">
                            <div class="mb-3">
                                <label for="cst_name" class="form-label">Name</label>
                                <input type="text" class="form-control" id="cst_name" name="cst_name" required>
                            </div>
                            <div class="mb-3">
                                <label for="cst_dob" class="form-label">DOB</label>
                                <input type="date" class="form-control" id="cst_dob" name="cst_dob" required>
                            </div>
                            <div class="mb-3">
                                <label for="cst_email" class="form-label">Email</label>
                                <input type="email" class="form-control" id="cst_email" name="cst_email" required>
                            </div>
                            <div class="mb-3">
                                <label for="cst_phoneNum" class="form-label">Phone</label>
                                <input type="text" class="form-control" id="cst_phoneNum" name="cst_phoneNum" required>
                            </div>
                            <div class="mb-3">
                                <label for="nationality_id" class="form-label">Nationality</label>
                                <select class="form-control" id="nationality_id" name="nationality_id" required>
                                    <option value="">Select Nationality</option>
                                    @foreach ($nationalities as $nationality)
                                        <option value="{{ $nationality['nationality_id'] }}">{{ $nationality['nationality_name'] }}</option>
                                    @endforeach
                                </select>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <h5>Family</h5>
                            <div id="family-list-container">
                            </div>
                            <button type="button" class="btn btn-secondary mt-3" id="add-family">
                                + Add
                            </button>
                            <div id="family-error" class="text-danger mt-2" style="display: none;">Please add at least one family member.</div>
                        </div>
                    </div>
                    <button type="submit" class="btn btn-primary mt-4">Save</button>
                </form>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3>Customer List</h3>
            </div>
            <div class="card-body">
                <div class="table-responsive">
                    <table class="table table-striped">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Name</th>
                                <th>DOB</th>
                                <th>Email</th>
                                <th>Number of Families</th>
                                <th>Action</th>
                            </tr>
                        </thead>
                        <tbody>
                            @forelse ($customers->data as $customer)
                                <tr>
                                    <td>{{ $customer['cst_id'] }}</td>
                                    <td>{{ $customer['cst_name'] }}</td>
                                    <td>{{ \Carbon\Carbon::parse($customer['cst_dob'])->format('d M Y') }}</td>
                                    <td>{{ $customer['cst_email'] }}</td>
                                    <td>{{ $customer['family_count'] }}</td>
                                    <td>
                                        <button type="button" class="btn btn-sm btn-info edit-customer-btn" 
                                                data-bs-toggle="modal" 
                                                data-bs-target="#editCustomerModal" 
                                                data-id="{{ $customer['cst_id'] }}">
                                            Edit
                                        </button>
                                        <button type="button" class="btn btn-sm btn-danger delete-customer-btn" data-id="{{ $customer['cst_id'] }}">
                                            Delete
                                        </button>
                                    </td>
                                </tr>
                            @empty
                                <tr>
                                    <td colspan="6" class="text-center">No customer data found</td>
                                </tr>
                            @endforelse
                        </tbody>
                    </table>
                </div>
            </div>
            <div class="card-footer">
                @if ($customers->total_data > $customers->limit)
                    <nav>
                        <ul class="pagination justify-content-center mb-0">
                            @if ($customers->page > 1)
                                <li class="page-item"><a class="page-link" href="{{ url('/?page=' . ($customers->page - 1) . '&limit=' . $customers->limit) }}">Previous</a></li>
                            @else
                                <li class="page-item disabled"><span class="page-link">Previous</span></li>
                            @endif
    
                            @php
                                $totalPages = ceil($customers->total_data / $customers->limit);
                                $startPage = max(1, $customers->page - 2);
                                $endPage = min($totalPages, $customers->page + 2);
                            @endphp
                            @for ($i = $startPage; $i <= $endPage; $i++)
                                <li class="page-item @if($i == $customers->page) active @endif"><a class="page-link" href="{{ url('/?page=' . $i . '&limit=' . $customers->limit) }}">{{ $i }}</a></li>
                            @endfor
                            
                            @if ($customers->page < $totalPages)
                                <li class="page-item"><a class="page-link" href="{{ url('/?page=' . ($customers->page + 1) . '&limit=' . $customers->limit) }}">Next</a></li>
                            @else
                                <li class="page-item disabled"><span class="page-link">Next</span></li>
                            @endif
                        </ul>
                    </nav>
                @endif
            </div>
        </div>
    </div>

    <div class="modal fade" id="editCustomerModal" tabindex="-1" aria-labelledby="editCustomerModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-xl">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="editCustomerModalLabel">Edit Customer</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <form id="editCustomerForm" method="POST">
                    @csrf
                    @method('PUT')
                    <input type="hidden" name="cst_id" id="edit-cst-id">
                    <div class="modal-body">
                        <div class="row">
                            <div class="col-md-6">
                                <div class="mb-3">
                                    <label for="edit-cst-name" class="form-label">Name</label>
                                    <input type="text" class="form-control" id="edit-cst-name" name="cst_name" required>
                                </div>
                                <div class="mb-3">
                                    <label for="edit-cst-dob" class="form-label">DOB</label>
                                    <input type="date" class="form-control" id="edit-cst-dob" name="cst_dob" required>
                                </div>
                                <div class="mb-3">
                                    <label for="edit-cst-email" class="form-label">Email</label>
                                    <input type="email" class="form-control" id="edit-cst-email" name="cst_email" required>
                                </div>
                                <div class="mb-3">
                                    <label for="edit-cst-phoneNum" class="form-label">Phone</label>
                                    <input type="text" class="form-control" id="edit-cst-phoneNum" name="cst_phoneNum" required>
                                </div>
                                <div class="mb-3">
                                    <label for="edit-nationality-id" class="form-label">Nationality</label>
                                    <select class="form-control" id="edit-nationality-id" name="nationality_id" required>
                                        @foreach ($nationalities as $nationality)
                                            <option value="{{ $nationality['nationality_id'] }}">{{ $nationality['nationality_name'] }}</option>
                                        @endforeach
                                    </select>
                                </div>
                            </div>
                            <div class="col-md-6">
                                <h5>Family</h5>
                                <div id="edit-family-list-container">
                                </div>
                                <button type="button" class="btn btn-secondary mt-3" id="edit-add-family">
                                    + Add
                                </button>
                                <div id="edit-family-error" class="text-danger mt-2" style="display: none;">Please add at least one family member.</div>
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                        <button type="submit" class="btn btn-primary">Save</button>
                    </div>
                </form>
            </div>
        </div>
    </div>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"></script>
    
    <script>
        const goApiUrl = "{{ env('GO_API_URL') }}";

        document.addEventListener('DOMContentLoaded', function() {
            function addFamilyForm(container, family = null) {
                const newFamilyForm = document.createElement('div');
                newFamilyForm.classList.add('row', 'mb-3', 'align-items-end', 'family-form-group');
                newFamilyForm.innerHTML = `
                    <div class="col-md-4">
                        <label for="family-name" class="form-label">Name</label>
                        <input type="text" class="form-control" name="family_name[]" value="${family ? family.fl_name : ''}" required>
                    </div>
                    <div class="col-md-3">
                        <label for="family-relation" class="form-label">Relation</label>
                        <input type="text" class="form-control" name="family_relation[]" value="${family ? family.fl_relation : ''}" required>
                    </div>
                    <div class="col-md-3">
                        <label for="family-dob" class="form-label">DOB</label>
                        <input type="date" class="form-control" name="family_dob[]" value="${family ? family.fl_dob : ''}" required>
                    </div>
                    <div class="col-md-1">
                        <button type="button" class="btn btn-danger remove-family">Delete</button>
                    </div>
                `;
                container.appendChild(newFamilyForm);
            }
            
            const createCustomerForm = document.getElementById('createCustomerForm');
            const createFamilyContainer = document.getElementById('family-list-container');
            const familyError = document.getElementById('family-error');
            const addFamilyBtn = document.getElementById('add-family');
            const customerListTable = document.querySelector('table tbody');
            
            if (createCustomerForm && createFamilyContainer && addFamilyBtn) {
                addFamilyBtn.addEventListener('click', function() {
                    addFamilyForm(createFamilyContainer);
                });
            
                createFamilyContainer.addEventListener('click', function(event) {
                    if (event.target.classList.contains('remove-family')) {
                        event.target.closest('.family-form-group').remove();
                        if (createFamilyContainer.children.length === 0) {
                            familyError.style.display = 'block';
                        }
                    }
                });

                createCustomerForm.addEventListener('submit', function(event) {
                    if (createFamilyContainer.children.length === 0) {
                        event.preventDefault();
                        familyError.style.display = 'block';
                    } else {
                        familyError.style.display = 'none';
                    }
                });
            }

            const editModal = document.getElementById('editCustomerModal');
            const form = document.getElementById('editCustomerForm');
            const editFamilyContainer = document.getElementById('edit-family-list-container');
            const editFamilyError = document.getElementById('edit-family-error');
            const editAddFamilyBtn = document.getElementById('edit-add-family');

            if (editModal && form && editFamilyContainer && editAddFamilyBtn) {
                editModal.addEventListener('show.bs.modal', function(event) {
                    const button = event.relatedTarget;
                    const customerId = button.getAttribute('data-id');

                    form.action = '/customers/' + customerId;
                    form.reset();
                    editFamilyContainer.innerHTML = '';
                    
                    fetch(goApiUrl + '/customers/' + customerId)
                        .then(response => {
                            if (!response.ok) {
                                throw new Error('Network response was not ok');
                            }
                            return response.json();
                        })
                        .then(data => {
                            if (data.status === 200) {
                                const customer = data.data.customer;
                                const familyLists = data.data.family_list;
            
                                document.getElementById('edit-cst-id').value = customer.cst_id;
                                document.getElementById('edit-cst-name').value = customer.cst_name;
                                document.getElementById('edit-cst-dob').value = customer.cst_dob.split('T')[0];
                                document.getElementById('edit-cst-email').value = customer.cst_email;
                                document.getElementById('edit-cst-phoneNum').value = customer.cst_phoneNum;
                                document.getElementById('edit-nationality-id').value = customer.nationality_id;
                                
                                familyLists.forEach(family => {
                                     addFamilyForm(editFamilyContainer, family);
                                });
                            } else {
                                alert('Failed to retrieve customer data.');
                            }
                        })
                        .catch(error => {
                            console.error('Error get data:', error);
                            alert('Error occurred when get data.');
                        });
                });
            
                editAddFamilyBtn.addEventListener('click', function() {
                    addFamilyForm(editFamilyContainer);
                });

                editFamilyContainer.addEventListener('click', function(event) {
                    if (event.target.classList.contains('remove-family')) {
                        event.target.closest('.family-form-group').remove();
                        if (editFamilyContainer.children.length === 0) editFamilyError.style.display = 'block';
                    }
                });

                form.addEventListener('submit', function(event) {
                    if (editFamilyContainer.children.length === 0) {
                        event.preventDefault();
                        editFamilyError.style.display = 'block';
                    } else {
                        editFamilyError.style.display = 'none';
                    }
                });
            }

            customerListTable.addEventListener('click', function(event) {
                if (event.target.classList.contains('delete-customer-btn')) {
                    const customerId = event.target.getAttribute('data-id');

                    if (confirm('Are you sure you want to delete this customer?')) {
                        fetch(`${goApiUrl}/customers/${customerId}`, {
                            method: 'DELETE',
                        })
                        .then(data => {
                            if (data.status === 200) {
                                alert('Customer successfully deleted');
                                window.location.reload();
                            } else {
                                alert('Failed to delete customer: ' + data.message);
                            }
                        })
                        .catch(error => {
                            console.error('Error:', error);
                            alert('Error occurred when deleting the customer');
                        });
                    }
                }
            });
        });
    </script>
</body>
</html>