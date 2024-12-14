<?php

namespace App\Http\Controllers;

use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\File;

class ImageController extends Controller
{
    public function index(Request $request): JsonResponse
    {
        $category = $request->input('category');
        $basePath = public_path("images");

        // If category is provided, validate it
        if ($category && !preg_match('/^[a-zA-Z0-9_\-]+$/', $category)) {
            return response()->json(['error' => 'Invalid category'], 400);
        }

        $path = $category ? "$basePath/$category" : $basePath;

        // Check if the directory exists
        if (!is_dir($path)) {
            return response()->json(['error' => 'Category not found'], 404);
        }

        // Get all files recursively if category is null, otherwise limit to the specific category
        $files = File::allFiles($path);

        // Generate URLs for all found files
        $imageUrls = array_map(function ($file) use ($category) {
            // Generate the relative URL for each image
            $relativePath = isset($category) ? "images/$category/{$file->getRelativePath()}" : "images/{$file->getRelativePath()}";
            return url($relativePath . $file->getFilename());
        }, $files);

        return response()->json($imageUrls);
    }
}
