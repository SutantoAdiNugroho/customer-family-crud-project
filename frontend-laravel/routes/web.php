<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\CustomerController;

Route::get('/', [CustomerController::class, 'index']);
Route::post('/', [CustomerController::class, 'store']);
Route::put('/customers/{id}', [CustomerController::class, 'update']);