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
                <form action="/" method="POST">
                    @csrf
                    <div class="row">
                        <div class="col-md-6">
                            <div class="mb-3">
                                <label for="cst_name" class="form-label">Name</label>
                                <input type="text" class="form-control" id="cst_name" name="cst_name" required>
                            </div>
                            <div class="mb-3">
                                <label for="cst_dob" class="form-label">Date of Birthday</label>
                                <input type="date" class="form-control" id="cst_dob" name="cst_dob" required>
                            </div>
                            <div class="mb-3">
                                <label for="cst_email" class="form-label">Email</label>
                                <input type="email" class="form-control" id="cst_email" name="cst_email" required>
                            </div>
                            <div class="mb-3">
                                <label for="cst_phoneNum" class="form-label">Phone Number</label>
                                <input type="text" class="form-control" id="cst_phoneNum" name="cst_phoneNum" required>
                            </div>
                            <div class="mb-3">
                                <label for="nationality_id" class="form-label">Nalionality</label>
                                <select class="form-control" id="nationality_id" name="nationality_id" required>
                                    <option value="">Select Nationality</option>
                                    @foreach ($nationalities as $nationality)
                                        <option value="{{ $nationality['nationality_id'] }}">{{ $nationality['nationality_name'] }}</option>
                                    @endforeach
                                </select>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <h5>Family Lists</h5>
                            <div id="family-list-container">
                                </div>
                            <button type="button" class="btn btn-secondary mt-3" id="add-family">
                                + Add
                            </button>
                        </div>
                    </div>
                    <button type="submit" class="btn btn-primary mt-4">Save</button>
                </form>
            </div>
        </div>

        <div class="card">
            <div class="card-header">
                <h3>List of Customers</h3>
            </div>
            <div class="card-body">
                <div class="table-responsive">
                    <table class="table table-striped">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Name</th>
                                <th>Date of Birthday</th>
                                <th>Email</th>
                                <th>Number of families</th>
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
                                        <a href="#" class="btn btn-sm btn-info">Edit</a>
                                        <a href="#" class="btn btn-sm btn-danger">Delete</a>
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

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"></script>
    <script>
        document.getElementById('add-family').addEventListener('click', function() {
            const container = document.getElementById('family-list-container');
            const familyCount = container.children.length;

            const newFamilyForm = document.createElement('div');
            newFamilyForm.classList.add('row', 'mb-3', 'align-items-end', 'family-form-group');
            newFamilyForm.innerHTML = `
                <div class="col-md-4">
                    <label for="family_name_${familyCount}" class="form-label">Name</label>
                    <input type="text" class="form-control" name="family_name[]" id="family_name_${familyCount}">
                </div>
                <div class="col-md-4">
                    <label for="family_relation_${familyCount}" class="form-label">Relation</label>
                    <input type="text" class="form-control" name="family_relation[]" id="family_relation_${familyCount}">
                </div>
                <div class="col-md-3">
                    <label for="family_dob_${familyCount}" class="form-label">Date of Birthday</label>
                    <input type="date" class="form-control" name="family_dob[]" id="family_dob_${familyCount}">
                </div>
                <div class="col-md-1">
                    <button type="button" class="btn btn-danger remove-family">Delete</button>
                </div>
            `;
            container.appendChild(newFamilyForm);

            newFamilyForm.querySelector('.remove-family').addEventListener('click', function() {
                newFamilyForm.remove();
            });
        });

        document.querySelectorAll('.remove-family').forEach(button => {
            button.addEventListener('click', function() {
                button.closest('.family-form-group').remove();
            });
        });
    </script>
</body>
</html>