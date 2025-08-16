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
        $response = Http::post($this->goApiUrl . '/customers', [
            'customer' => [
                'nationality_id' => (int) $request->input('nationality_id'),
                'cst_name' => $request->input('cst_name'),
                'cst_dob' => Carbon::parse($request->input('cst_dob'))->toIso8601String(),
                'cst_phoneNum' => $request->input('cst_phoneNum'),
                'cst_email' => $request->input('cst_email'),
            ],
            'family_list' => $this->formatFamilyLists($request),
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