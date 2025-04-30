from django.urls import URLPattern, path
from .views import insert_post,show_name

urlpatterns=[
    path('create/',insert_post),
    path('show/',show_name)
]
