<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Validation\ValidationException;
use Carbon\Carbon;

class CustomerController extends Controller
{
    private $goApiUrl;

    public function __construct()
    {
        $this->goApiUrl = env('GO_API_URL');
    }

    public function index(Request $request)
    {
        $page = $request->query('page', 1);
        $limit = 10;

        // get nationalities
        $nationalities = [];
        $nationalityResponse = Http::get($this->goApiUrl . '/nationalities');
        if ($nationalityResponse->successful()) {
            $nationalities = $nationalityResponse->json()['data'] ?? [];
        }

        // get customer data with pagination
        $customerResponse = Http::get($this->goApiUrl . '/customers', [
            'page' => $page,
            'limit' => $limit,
        ]);

        $customers = (object) [];
        if ($customerResponse->successful()) {
            $customers = (object) $customerResponse->json()['data'] ?? (object) [];
        }

        return view('customers.index', compact('nationalities', 'customers'));
    }

    public function store(Request $request)
    {
        // validation for customer data
        $validatedCustomerData = $request->validate([
            'cst_name' => 'required|string|max:255',
            'cst_dob' => 'required|date',
            'cst_phoneNum' => 'required|string|max:20',
            'cst_email' => 'required|email|max:255',
            'nationality_id' => 'required|integer',
        ], [
            'cst_name.required' => 'Customer name is required',
            'cst_dob.required' => 'Date of birthday customer is required',
            'cst_phoneNum.required' => 'Phone number is required',
            'cst_email.required' => 'Email is required',
            'cst_email.email' => 'Email format is not valid',
            'nationality_id.required' => 'Customer nationality is required',
        ]);

        // validation for customer family list data
        $validatedFamilyData = $request->validate([
            'family_name.*' => 'required|string|max:255',
            'family_relation.*' => 'required|string|max:255',
            'family_dob.*' => 'required|date',
        ], [
            'family_name.*.required' => 'Family member names are required',
            'family_relation.*.required' => 'Family member relation are required',
            'family_dob.*.required' => 'Family member date of birthday are required',
        ]);
        
        $customerData = array_merge($validatedCustomerData, [
            'cst_dob' => Carbon::parse($validatedCustomerData['cst_dob'])->toIso8601String(),
            'nationality_id' => (int) $validatedCustomerData['nationality_id'],
        ]);
        
        $familyLists = [];
        if (isset($validatedFamilyData['family_name'])) {
            foreach ($validatedFamilyData['family_name'] as $key => $name) {
                $familyLists[] = [
                    'fl_name' => $name,
                    'fl_relation' => $validatedFamilyData['family_relation'][$key],
                    'fl_dob' => $validatedFamilyData['family_dob'][$key],
                ];
            }
        }

        $response = Http::post($this->goApiUrl . '/customers', [
            'customer' => $customerData,
            'family_list' => $familyLists,
        ]);

        if ($response->successful()) {
            return redirect('/')->with('success', 'Customer has been successfully added');
        }

        // error api handler from backend
        $errorData = $response->json();
        $message = $errorData['message'] ?? 'Error occured when add customer.';

        // redirect with error
        throw ValidationException::withMessages([
            'api_error' => [$message],
        ]);
    }

    public function update(Request $request, $id)
    {
        $validatedCustomerData = $request->validate([
            'cst_name' => 'required|string|max:255',
            'cst_dob' => 'required|date',
            'cst_phoneNum' => 'required|string|max:20',
            'cst_email' => 'required|email|max:255',
            'nationality_id' => 'required|integer',
        ]);
        
        $validatedFamilyData = $request->validate([
            'family_name.*' => 'required|string|max:255',
            'family_relation.*' => 'required|string|max:255',
            'family_dob.*' => 'required|date',
        ]);

        $customerData = array_merge($validatedCustomerData, [
            'cst_dob' => Carbon::parse($validatedCustomerData['cst_dob'])->toIso8601String(),
            'nationality_id' => (int) $validatedCustomerData['nationality_id'],
        ]);

        $familyLists = [];
        if (isset($validatedFamilyData['family_name'])) {
            foreach ($validatedFamilyData['family_name'] as $key => $name) {
                $familyLists[] = [
                    'fl_name' => $name,
                    'fl_relation' => $validatedFamilyData['family_relation'][$key],
                    'fl_dob' => $validatedFamilyData['family_dob'][$key],
                ];
            }
        }

        $response = Http::put($this->goApiUrl . '/customers/' . $id, [
            'customer' => $customerData,
            'family_list' => $familyLists,
        ]);

        if ($response->successful()) {
            return redirect('/')->with('success', 'Customer successfully updated');
        }

        $errorData = $response->json();
        $message = $errorData['message'] ?? 'Error occured when update customer';

        throw ValidationException::withMessages([
            'api_error' => [$message],
        ]);
    }

    private function formatFamilyLists(Request $request)
    {
        $familyLists = [];
        $familyNames = $request->input('family_name');
        $familyRelations = $request->input('family_relation');
        $familyDobs = $request->input('family_dob');

        if ($familyNames && is_array($familyNames)) {
            foreach ($familyNames as $key => $name) {
                if (!empty($name)) {
                    $familyLists[] = [
                        'fl_relation' => $familyRelations[$key],
                        'fl_name' => $name,
                        'fl_dob' => $familyDobs[$key],
                    ];
                }
            }
        }
        return $familyLists;
    }
}