from rest_framework.decorators import api_view
from rest_framework.response import Response
from rest_framework import status 
from django.http import HttpRequest,HttpResponse
from django.shortcuts import render
from .serializers import PostSerializers 

# Create your views here.
@api_view(["POST"])
def insert_post(request):
    print(f"hello is there anything==>{request.data}")
    serializer=PostSerializers(data=request.data) 
    if serializer.is_valid():
        serializer.save()
        return Response(serializer.data,status=status.HTTP_201_CREATED)
    return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

@api_view(["GET"])
def show_name(request):
    json={
        "name":"kushal",
        "bimohit":"shrestha",
        "office":'codequant'
    }
    return Response(json,status=status.HTTP_201_CREATED)
